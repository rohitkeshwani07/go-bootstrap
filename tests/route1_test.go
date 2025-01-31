package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rohitkeshwani07/go-bootstrap/internal/users"
	"github.com/rohitkeshwani07/go-bootstrap/mocks"
	"github.com/rohitkeshwani07/go-bootstrap/pkg/router"
	"github.com/rohitkeshwani07/go-bootstrap/routes"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/mock/gomock"
)

func CreateTestApp(overrides ...fx.Option) (*gin.Engine, func()) {
	var r *gin.Engine

	// Create an fx app with both real and override providers
	app := fx.New(
		fx.Provide(
			router.NewRouter,
			routes.NewRoutes,
			users.NewUserService,
			fx.Annotate(
				users.NewBusinessLogic,
				fx.As(new(users.IBusinessLogic)),
			),
		),
		fx.Options(overrides...), // Apply any overrides
		fx.Populate(&r),          // Populate router from fx container
	)

	// Start the app before returning
	err := app.Start(context.Background())
	if err != nil {
		panic(err) // Fail fast in case of setup issues
	}

	// Return the router and a cleanup function
	return r, func() {
		app.Stop(context.Background()) // Ensure proper shutdown after test
	}
}

func TestFunctionA(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockBusinessLogic := mocks.NewMockIBusinessLogic(ctrl)
	mockBusinessLogic.EXPECT().GetUser().Return("Rohit").Times(1)

	router, cleanup := CreateTestApp(
		fx.Replace(fx.Annotate(
			mockBusinessLogic,
			fx.As(new(users.IBusinessLogic)),
		)),
	)
	defer cleanup()

	// Create a test request
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check response body
	expectedBody := `{"message":"Rohit"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
