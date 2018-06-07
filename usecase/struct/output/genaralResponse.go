package output

type HttpHeader struct {
	StatusCode int
}

type ErrorBody struct {
	ErrorMessage string `json:"error_message"`
}
