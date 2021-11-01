package helpers

import (
	"encoding/json"
	"net/http"
	"saiga/models"
)

func HTTPErrorHandler(w http.ResponseWriter, body string, status int) {
	w.WriteHeader(status)
	errorResponse := models.HttpErrorResponse{
		Body: body,
	}
	jData, _ := json.Marshal(errorResponse)
	w.Write(jData)
}
