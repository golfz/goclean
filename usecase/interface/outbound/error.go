package outbound

import "goclean/usecase/struct/output"

type ErrorResponse interface {
	PresentErrorResponse(header output.HttpHeader, body output.ErrorBody)
}
