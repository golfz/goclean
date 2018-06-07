package inbound

import (
	"goclean/usecase/struct/input"
)

type ExampleRequest interface {
	GetExampleData(data input.Example)
}
