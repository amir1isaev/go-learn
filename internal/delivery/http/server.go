package http

import (
	"context"
	"fmt"
	"go-learn/internal/config"
	"go-learn/internal/delivery/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(router *gin.Engine, handlers *handlers.Handlers) *Server {
	httpCfg := config.Get().Transport.HTTP
	addr := fmt.Sprintf("%s:%s", httpCfg.Host, httpCfg.Port)

	handlers.MapHandlers(router)

	return &Server{
		httpServer: &http.Server{
			Addr:           addr,
			Handler:        router.Handler(),
			MaxHeaderBytes: httpCfg.MaxHeaderMegabytes,
			ReadTimeout:    httpCfg.Timeouts.Read,
			WriteTimeout:   httpCfg.Timeouts.Write,
		},
	}
}

func (s Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
