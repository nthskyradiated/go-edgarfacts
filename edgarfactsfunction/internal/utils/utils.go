package utils

import (
	"fmt"
	"net/http"
)

func HandleHttpErr(w http.ResponseWriter, msg string, err error, statusCode int) {
	w.WriteHeader(statusCode)
	if err != nil {
		msg = fmt.Sprintf("%s: %v", msg, err)
	}
	fmt.Fprint(w, msg)
} 