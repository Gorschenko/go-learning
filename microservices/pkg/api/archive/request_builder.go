package api

type RequestBuilder[Body any] struct {
	URL    string
	Method HTTPMethod
	Body   Body
	Query  map[string]string
}

func (builder *RequestBuilder[Body]) SetMethod(method HTTPMethod) *RequestBuilder[Body] {
	builder.Method = method

	return builder
}

func (builder *RequestBuilder[Body]) SetURL(url string) *RequestBuilder[Body] {
	builder.URL = url
	return builder
}

func (builder *RequestBuilder[Body]) SetBody(body Body) *RequestBuilder[Body] {
	builder.Body = body

	return builder
}

func (builder *RequestBuilder[Body]) SetQuery(key, value string) *RequestBuilder[Body] {
	if builder.Query == nil {
		builder.Query = make(map[string]string)
	}

	builder.Query[key] = value

	return builder
}
