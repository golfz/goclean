package outbound

import "goclean/usecase/struct/output"

type ExampleResponse interface {
	PresentExampleResult(header output.HttpHeader, data output.Example)
}
