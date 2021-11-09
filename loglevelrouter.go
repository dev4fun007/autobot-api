package api

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"strings"
)

func LogLevelRouteHandler(r *mux.Router) *mux.Router {
	r.HandleFunc("/loglevel", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		_, _ = writer.Write([]byte(zerolog.GlobalLevel().String()))
	}).Methods("GET")

	r.HandleFunc("/loglevel/{level}", func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		level := strings.ToLower(params["level"])
		switch level {
		case "panic":
			zerolog.SetGlobalLevel(zerolog.PanicLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		case "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "warn":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "info":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "trace":
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	}).Methods("PUT")
	return r
}
