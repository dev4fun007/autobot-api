package api

import (
	"context"
	"fmt"
	"github.com/dev4fun007/autobot-common"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

const (
	ServerTag = "ApiServer"
)

type ConfigApiServer struct {
	*http.Server
}

func NewApiServer(port string, r *mux.Router) ConfigApiServer {
	return ConfigApiServer{
		Server: &http.Server{
			Addr:         fmt.Sprintf("0.0.0.0:%s", port),
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      r,
		},
	}
}

func (apiServer ConfigApiServer) StartApiServer(ctx context.Context) {
	go func() {
		log.Info().Str(common.LogComponent, ServerTag).Msg("starting config server...")
		if err := apiServer.Server.ListenAndServe(); err != nil {
			log.Error().Str(common.LogComponent, ServerTag).Err(err).Msg("config server shutting down...")
		}
	}()
}

func (apiServer ConfigApiServer) ShutdownApiServer(ctx context.Context) {
	_ = apiServer.Server.Shutdown(ctx)
}
