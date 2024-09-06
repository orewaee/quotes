package utils

import (
	"github.com/orewaee/quotes/internal/dtos"
	"net/http"
)

func MustWriteError(writer http.ResponseWriter, message string, code int) {
	data := &dtos.Error{
		Message: message,
	}

	MustWriteJson(writer, data, code)
}
