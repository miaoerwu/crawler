package collect

type Request struct {
	Url       string
	Cookies   string
	ParseFunc func([]byte, *Request) ParseResult
}

type ParseResult struct {
	Requests []*Request
	Items    []any
}
