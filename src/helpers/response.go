package helpers

import (
	"encoding/json"
	"goauth/v2/src/models"
	"net/http"
)

type ResponseType struct {
	writer http.ResponseWriter
}

func Response(w http.ResponseWriter) ResponseType {
	return ResponseType{w}
}

func (resp ResponseType) Write(data models.ResponseObject, code int) {
	data.Status = code
	resp.writer.Header().Add("Content-Type", "application/json")
	resp.writer.WriteHeader(code)
	json.NewEncoder(resp.writer).Encode(data)
}
