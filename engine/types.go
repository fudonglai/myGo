package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items    []interface{} // data
}

func NilParser(url []byte) ParserResult  {
	return ParserResult{}
}
