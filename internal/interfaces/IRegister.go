package interfaces

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type IRegisterService interface {
	Register(c context.Context, req models.User) (interface{}, error)
}

type IRegisterHandler interface {
	Register(c *gin.Context)
}
