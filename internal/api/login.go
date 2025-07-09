package api

import (
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService interfaces.ILoginService
}

func (api *LoginHandler) Login(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.LoginRequest
		res models.LoginResponse
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailledBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailledBadRequest, nil)
		return
	}

	res, err := api.LoginService.Login(c.Request.Context(), req)
	if err != nil {
		log.Error("failed on login service: ", err)
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHttp(c, http.StatusOK, constants.SuccessMessage, res)
}
