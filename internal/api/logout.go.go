package api

import (
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	LogoutService interfaces.ILogoutService
}

func (api *LogoutHandler) Logout(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	token := c.Request.Header.Get("Authorization")
	err := api.LogoutService.Logout(c.Request.Context(), token)
	if err != nil {
		log.Error("failed on logut service: ", err)
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHttp(c, http.StatusOK, constants.SuccessMessage, nil)
}
