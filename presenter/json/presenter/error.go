package presenter

import (
	"goclean/presenter/json/viewinterface"
	"goclean/usecase/struct/output"
)

type errorPresenter struct {
	view viewinterface.RestfulView
}

func InitiateErrorPresenter(view viewinterface.RestfulView) *errorPresenter {
	return &errorPresenter{view: view}
}

func (p *errorPresenter) PresentErrorResponse(header output.HttpHeader, body output.ErrorBody) {
	p.view.Write(header, body)
}
