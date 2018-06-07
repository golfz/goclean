package controller

import (
	"net/http"
	"goclean/usecase/interface/outbound"
	"goclean/usecase/interface/inbound"
	"goclean/usecase/struct/input"
	"goclean/usecase/struct/output"
	"github.com/gorilla/mux"
)

type exampleCtrl struct {
}

func InitiateExampleCtrl() *exampleCtrl {
	return &exampleCtrl{}
}

// GET Example
func (ctrl *exampleCtrl) GetExampleData(r *http.Request, uc inbound.ExampleRequest, pError outbound.ErrorResponse) {
	vars := mux.Vars(r)
	key, hasKey := vars["key"]

	if hasKey == false || key == "" {
		header := output.HttpHeader{StatusCode: http.StatusBadRequest}
		errBody := output.ErrorBody{ErrorMessage: "key not found"}
		pError.PresentErrorResponse(header, errBody)
	}

	inputData := input.Example{
		Key:key,
	}

	uc.GetExampleData(inputData)
}
