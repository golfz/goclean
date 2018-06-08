package interactor

import (
	"goclean/usecase/interface/daointerface"
	"goclean/usecase/interface/outbound"
	"net/http"
	"goclean/usecase/struct/input"
	"goclean/usecase/struct/output"
)

type exampleUseCase struct {
	dao      daointerface.ExampleDaoInterface
	pSuccess outbound.ExampleResponse
	pError   outbound.ErrorResponse
}

func InitiateExampleUseCase(dao daointerface.ExampleDaoInterface, pSuccess outbound.ExampleResponse, pError outbound.ErrorResponse) *exampleUseCase {
	return &exampleUseCase{dao: dao, pSuccess: pSuccess, pError: pError}
}

func (uc *exampleUseCase) GetExampleData(data input.Example) {
	key := data.Key
	r := uc.dao.GetExampleData(key)
	responseBody := output.Example{
		Key:         r.Key,
		Name:        r.Name,
		Color:       r.Color,
		Remain:      r.Remain,
		CreatedDate: r.CreatedDate,
	}

	uc.pSuccess.PresentExampleResult(
		output.HttpHeader{
			StatusCode: http.StatusOK,
		},
		responseBody,
	)

}
