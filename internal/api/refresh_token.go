package api

import (
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRefreshTokenService
}

func (api *RefreshTokenHandler) RefreshToken(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	refreshToken := c.Request.Header.Get("Authorization")
	claim, ok := c.Get("token")
	if !ok {
		log.Error("failed to get claim token in context")
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrServerError, nil)
		return
	}

	//assertion (mengambil dar)
	tokenClaim, ok := claim.(*helpers.ClaimToken)
	if !ok {
		log.Error("failed to parse claim token")
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrServerError, nil)
		return
	}

	resp, err := api.RefreshTokenService.RefreshToken(c.Request.Context(), refreshToken, *tokenClaim)
	if err != nil {
		log.Error("failed on logut service: ", err)
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHttp(c, http.StatusOK, constants.SuccessMessage, resp)
}
