package presenter

import (
	"goclean/usecase/struct/output"
	"goclean/presenter/json/viewstruct"
	"goclean/presenter/json/viewinterface"
	"fmt"
)

type examplePresenter struct {
	view viewinterface.RestfulView
}

func InitiateExamplePresenter(view viewinterface.RestfulView) *examplePresenter {
	return &examplePresenter{view: view}
}

func (p *examplePresenter) PresentExampleResult(header output.HttpHeader, data output.Example) {
	url := fmt.Sprintf("/example/%s", data.Key)
	createDate := data.CreatedDate.Format("2006-01-02")
	body := viewstruct.Example{
		Key:         data.Key,
		Name:        data.Name,
		Color:       data.Color,
		Remain:      data.Remain,
		CreatedDate: createDate,
		Url:         url,
		Rel:         "self",
	}
	p.view.Write(header, body)
}
