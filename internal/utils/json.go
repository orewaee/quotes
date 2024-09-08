package utils

import "net/http"

func MustWriteJson(writer http.ResponseWriter, data interface{}, code int) {
	writer.Header().Set("Content-Type", "application/json")
	MustWriteAny(writer, data, code)
}
