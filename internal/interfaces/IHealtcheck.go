package interfaces

import "github.com/gin-gonic/gin"

type IHealthcheckServices interface {
	HealtcheckService() (string, error)
}
type IHealthcheckHandler interface {
	HealtcheckHandlerHTTP(c *gin.Context)
}

type IHealthcheckRepository interface {
}
