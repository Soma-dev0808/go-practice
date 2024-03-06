package apperrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError
	// 型変換 error => MyAppError
	if !errors.As(err, &appErr) {
		fmt.Println(err)
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err: err,
		}
	}

	var statusCode int
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}