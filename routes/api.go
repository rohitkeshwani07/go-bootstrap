package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rohitkeshwani07/go-bootstrap/internal/users"
	"github.com/rohitkeshwani07/go-bootstrap/pkg/router"
)

type Routes struct {
	UserService *users.UserService
}

// TODO: Middleware?
func (r *Routes) RegisterRoutes(router gin.IRouter) {
	router.GET("/ping", r.UserService.GetUser)
}

func NewRoutes(u *users.UserService) router.IRoutes {
	return &Routes{
		UserService: u,
	}
}
