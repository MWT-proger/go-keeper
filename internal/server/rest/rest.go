package rest

import (
	"net/http"

	"github.com/MWT-proger/go-keeper/configs/server/config"
	"github.com/MWT-proger/go-keeper/internal/pkg/logger"
	"github.com/MWT-proger/go-keeper/internal/server/handlers"
)

// Run() запускает сервер и слушает его по указанному хосту
func Run(h *handlers.APIHandler, conf *config.Config) error {
	r := router(h, conf)

	logger.Log.Info("Running server on", logger.StringField("host", conf.HostServer))

	return http.ListenAndServe(conf.HostServer, r)
}
