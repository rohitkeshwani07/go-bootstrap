package main

import (
	"time"

	"github.com/rohitkeshwani07/go-bootstrap/internal/users"
	"github.com/rohitkeshwani07/go-bootstrap/pkg/router"
	"github.com/rohitkeshwani07/go-bootstrap/routes"
	"go.uber.org/fx"
)

const MaxTimeGracefullShutdown = 60 * time.Second

func ProvideWithBinding(constructor interface{}, iface interface{}) fx.Option {
	return fx.Provide(
		fx.Annotate(
			constructor,
			fx.As(iface),
		),
	)
}

func main() {
	fx.New(
		fx.StopTimeout(MaxTimeGracefullShutdown),
		fx.Provide(
			router.NewRouter,
			routes.NewRoutes,
			users.NewUserService,
		),
		ProvideWithBinding(users.NewBusinessLogic, new(users.IBusinessLogic)),
		fx.Invoke(router.NewHTTPServer),
	).Run()
}
