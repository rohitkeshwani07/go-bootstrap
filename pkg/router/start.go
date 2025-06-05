package router

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohitkeshwani07/go-bootstrap/pkg/config"
	"go.uber.org/fx"
)

func NewRouter(routes IRoutes) *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r)
	return r
}

func NewHTTPServer(cfg *config.Config, lc fx.Lifecycle, r *gin.Engine) {
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server forced to shutdown: ", err)
			}
			return nil
		},
	})
}
