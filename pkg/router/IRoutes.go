package router

import "github.com/gin-gonic/gin"

type IRoutes interface {
	RegisterRoutes(router gin.IRouter)
}
