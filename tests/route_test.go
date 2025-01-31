package tests

import (
	"testing"

	"github.com/rohitkeshwani07/go-bootstrap/internal/users"
	"github.com/rohitkeshwani07/go-bootstrap/pkg/router"
	"github.com/rohitkeshwani07/go-bootstrap/routes"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestAppLifecycle(t *testing.T) {
	t.Skip("skipping as port is used by TestMain")

	app := fxtest.New(
		t,
		fx.Provide(
			routes.NewRoutes,
			users.NewUserService,
			fx.Annotate(
				users.NewBusinessLogic,
				fx.As(new(users.IBusinessLogic)),
			),
		),
		fx.Invoke(router.NewHTTPServer),
	)

	app.RequireStart()
	app.RequireStop()
}

// func TestMain(m *testing.M) {
// 	// Setup before running tests
// 	fmt.Println("Setup before tests")
// 	startCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	fx.New(
// 		fx.Provide(
// 			routes.NewRoutes,
// 			users.NewUserService,
// 			fx.Annotate(
// 				users.NewBusinessLogic,
// 				fx.As(new(users.IBusinessLogic)),
// 			),
// 		),
// 		fx.Invoke(router.NewHTTPServer),
// 	).Start(startCtx)

// 	// Run tests
// 	exitCode := m.Run()

// 	// Teardown after tests
// 	fmt.Println("Teardown after tests")

// 	// Exit with the appropriate code
// 	os.Exit(exitCode)
// }
