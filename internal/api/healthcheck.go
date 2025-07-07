package api

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Healthcheck struct {
	HealthcheckServices interfaces.IHealthcheckServices
}

func (api *Healthcheck) HealtcheckHandlerHTTP(c *gin.Context) {
	msg, err := api.HealthcheckServices.HealtcheckService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	helpers.SendResponseHttp(c, http.StatusOK, msg, nil)
}
