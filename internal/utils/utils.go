package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MustWriteBytes(writer http.ResponseWriter, data []byte, code int) {
	writer.WriteHeader(code)

	if _, err := writer.Write(data); err != nil {
		panic(err)
	}
}

func MustWriteString(writer http.ResponseWriter, data string, code int) {
	writer.WriteHeader(code)

	if _, err := fmt.Fprintln(writer, data); err != nil {
		panic(err)
	}
}

func MustWriteAny(writer http.ResponseWriter, data any, code int) {
	writer.WriteHeader(code)

	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	if _, err := writer.Write(bytes); err != nil {
		panic(err)
	}
}
