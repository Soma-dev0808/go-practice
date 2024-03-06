package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func SendSuccessJSONResponse(w http.ResponseWriter, data interface{}) {
	SendJSONResponse(w, data, http.StatusOK)
}

func SendSuccessPostJSONResponse(w http.ResponseWriter, data interface{}) {
	SendJSONResponse(w, data, http.StatusCreated)
}

type StringResponse struct{
	Result string `json:"result"`
}


func SendSuccessUpdateJSONResponse(w http.ResponseWriter) {
	response := StringResponse{
		Result: "updated",
	}
	SendJSONResponse(w, response, http.StatusOK)
}