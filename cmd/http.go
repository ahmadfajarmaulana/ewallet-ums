package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	healtcheckSvc := &services.Healcheck{}
	healcheckAPI := &api.Healthcheck{
		HealthcheckServices: healtcheckSvc,
	}

	r := gin.Default()

	r.GET("/healt", healcheckAPI.HealtcheckHandlerHTTP)

	err := r.Run(":" + helpers.GetEnv("APP_PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}
