package proxy

import (
	"errors"
	"net/http"
	"net/url"
	"sync/atomic"
)

type Func func(*http.Request) (*url.URL, error)

type roundRobinSwitcher struct {
	urls  []*url.URL
	index uint32
}

func NewRoundRobinSwitcher(proxyURLs ...string) (Func, error) {
	if len(proxyURLs) < 1 {
		return nil, errors.New("proxy URL list is empty")
	}

	urls := make([]*url.URL, len(proxyURLs))
	for i, v := range proxyURLs {
		u, err := url.Parse(v)
		if err != nil {
			return nil, err
		}
		urls[i] = u
	}

	return (&roundRobinSwitcher{urls: urls}).GetProxy, nil
}

func (r *roundRobinSwitcher) GetProxy(*http.Request) (*url.URL, error) {
	index := atomic.AddUint32(&r.index, 1) - 1

	u := r.urls[index%uint32(len(r.urls))]

	return u, nil
}
