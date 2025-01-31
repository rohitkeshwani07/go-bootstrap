package main

import (
	"time"

	"github.com/rohitkeshwani07/go-bootstrap/internal/users"
	"github.com/rohitkeshwani07/go-bootstrap/pkg/router"
	"github.com/rohitkeshwani07/go-bootstrap/routes"
	"go.uber.org/fx"
)

const MaxTimeGracefullShutdown = 60 * time.Second

/*
TODO: Create simple APIs for fx.Provide with Annotate.
Example:

	fxwrapper.ProvideNamedValueWithBinding("echo", NewEchoHandler, new(Route)),
*/
func main() {
	fx.New(
		fx.StopTimeout(MaxTimeGracefullShutdown),
		fx.Provide(
			router.NewRouter,
			routes.NewRoutes,
			users.NewUserService,
			fx.Annotate(
				users.NewBusinessLogic,
				fx.As(new(users.IBusinessLogic)),
			),
		),
		fx.Invoke(router.NewHTTPServer),
	).Run()
}
