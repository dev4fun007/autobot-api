package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	StrategyTypePathParam = "strategy_type"
	ConfigNamePathParam   = "config_name"
)

func NewConfigApiRouter(configApiHandler ConfigApiHandler) *mux.Router {
	r := mux.NewRouter()
	r = LogLevelRouteHandler(r)

	// register routes and handler
	s := r.PathPrefix("/config").Subrouter()
	s.HandleFunc(fmt.Sprintf("/strategy/{%s}", StrategyTypePathParam), configApiHandler.SaveConfig).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")

	s.HandleFunc(fmt.Sprintf("/strategy/{%s}/name/{%s}", StrategyTypePathParam, ConfigNamePathParam), configApiHandler.UpdateConfig).
		Methods(http.MethodPut).
		Headers("Content-Type", "application/json")

	s.HandleFunc(fmt.Sprintf("/strategy/{%s}/name/{%s}", StrategyTypePathParam, ConfigNamePathParam), configApiHandler.GetConfigByNameAndType).
		Methods(http.MethodGet)

	s.HandleFunc(fmt.Sprintf("/strategy/{%s}/name/{%s}", StrategyTypePathParam, ConfigNamePathParam), configApiHandler.DeleteConfigByNameAndType).
		Methods(http.MethodDelete)

	s.HandleFunc(fmt.Sprintf("/strategies/type/{%s}", StrategyTypePathParam), configApiHandler.GetConfigByType).
		Methods(http.MethodGet)

	return r
}
