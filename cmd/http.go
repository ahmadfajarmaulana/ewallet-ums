package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	dependency := dependencyInject()

	r := gin.Default()

	r.GET("/healt", dependency.HealtCheckApi.HealtcheckHandlerHTTP)

	userV1 := r.Group("/api/v1")
	userV1.POST("/register", dependency.RegisterApi.Register)
	userV1.POST("/login", dependency.LoginApi.Login)
	userV1.DELETE("/logout", dependency.MiddlewareValidateAuth, dependency.LogoutApi.Logout)
	userV1.PUT("/refresh-token", dependency.MiddlewareRefreshToken, dependency.RefreshTokenApi.RefreshToken)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	UserRepository interfaces.IUserRepository

	HealtCheckApi   interfaces.IHealthcheckHandler
	RegisterApi     interfaces.IRegisterHandler
	LoginApi        interfaces.ILoginHandler
	LogoutApi       interfaces.ILogoutHandler
	RefreshTokenApi interfaces.IRefreshTokenHandler

	//grpc
	TokenValidationApi *api.TokenValidationHandler
}

func dependencyInject() Dependency {
	healtcheckSvc := &services.Healcheck{}
	healcheckAPI := &api.Healthcheck{HealthcheckServices: healtcheckSvc}

	userRepo := &repository.UserRepository{DB: helpers.DB}
	registerservice := &services.RegisterService{UserRepository: userRepo}
	registerHandler := &api.RegisterHandler{RegisterService: registerservice}

	loginService := &services.LoginService{UserRepository: userRepo}
	loginHandler := &api.LoginHandler{LoginService: loginService}

	logoutService := &services.LogoutService{UserRepository: userRepo}
	logoutHandler := &api.LogoutHandler{LogoutService: logoutService}

	refreshTokenService := &services.RefreshTokenService{UserRepository: userRepo}
	refreshTokenHandler := &api.RefreshTokenHandler{RefreshTokenService: refreshTokenService}

	tokenValidationService := &services.TokenValidationService{UserRepository: userRepo}
	tokenValidationHandler := &api.TokenValidationHandler{
		TokenValidationService: tokenValidationService,
	}

	return Dependency{
		UserRepository: userRepo,

		HealtCheckApi:      healcheckAPI,
		RegisterApi:        registerHandler,
		LoginApi:           loginHandler,
		LogoutApi:          logoutHandler,
		RefreshTokenApi:    refreshTokenHandler,
		TokenValidationApi: tokenValidationHandler,
	}
}
