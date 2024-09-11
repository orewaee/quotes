package middlewares

import (
	"github.com/rs/zerolog"
	"net/http"
)

func LogMiddleware(next http.Handler, log *zerolog.Logger) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Debug().Str("path", request.URL.Path).Send()
		next.ServeHTTP(writer, request)
	})
}
