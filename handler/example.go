package handler

import (
	"net/http"
	"goclean/controller"
	"goclean/presenter/json/presenter"
	"goclean/dao"
	"goclean/usecase/interactor"
	"goclean/view/resful/json"
	"goclean/connector"
)

type exampleHandler struct {
}

func InitiateExampleHandler() *exampleHandler {
	return &exampleHandler{}
}

func (h *exampleHandler) GetExampleData(w http.ResponseWriter, r *http.Request) {

	dbSession := connector.InitiateMysqlConnectionSession()
	defer dbSession.Close()

	exampleDao := dao.InitiateExampleDao(dbSession)

	v := json.InitiateRestfulJsonView(w)

	pSuccess := presenter.InitiateExamplePresenter(v)
	pError := presenter.InitiateErrorPresenter(v)

	useCase := interactor.InitiateExampleUseCase(exampleDao, pSuccess, pError)

	ctrl := controller.InitiateExampleCtrl()

	ctrl.GetExampleData(r, useCase, pError)
}
