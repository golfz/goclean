package json

import (
	"net/http"
	"encoding/json"
	"goclean/usecase/struct/output"
)

type restfulJsonView struct {
	writer http.ResponseWriter
}

func InitiateRestfulJsonView(writer http.ResponseWriter) *restfulJsonView {
	return &restfulJsonView{writer: writer}
}

func (v *restfulJsonView) Write(header output.HttpHeader, body interface{}) {
	v.writer.Header().Set("Content-Type", "application/json")
	v.writer.WriteHeader(header.StatusCode)
	json.NewEncoder(v.writer).Encode(body)
}
