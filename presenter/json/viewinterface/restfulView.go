package viewinterface

import "goclean/usecase/struct/output"

type RestfulView interface {
	Write(header output.HttpHeader, body interface{})
}
