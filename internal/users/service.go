package users

import (
	"github.com/gin-gonic/gin"
)

type UserService struct {
	BusinessLogic IBusinessLogic
}

func NewUserService(businessLogic IBusinessLogic) *UserService {
	return &UserService{
		BusinessLogic: businessLogic,
	}
}

func (s *UserService) GetUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": s.BusinessLogic.GetUser(),
	})
}
