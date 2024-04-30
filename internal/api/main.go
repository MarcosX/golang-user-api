package api

import (
	"github.com/brizenox/golang-user-api/internal/domain"
	"github.com/brizenox/golang-user-api/internal/session"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupHandlers(echoInstance *echo.Echo) {
	userRepository := domain.NewUserRepository()

	echoInstance.Pre(middleware.RemoveTrailingSlash())
	echoInstance.Use(middleware.Recover())

	echoInstance.GET("/health", getHealth)

	userGroup := echoInstance.Group("/user", enforceValidSession())
	userHandler := NewUserHanlder(userRepository)
	userGroup.GET("", userHandler.getUser)
	userGroup.PUT("", userHandler.putUser)

	loginHandler := NewLoginHandler(userRepository)
	loginGroup := echoInstance.Group("/login")
	loginGroup.POST("", loginHandler.postLogin)

	signupHandler := NewSignupHandler(userRepository)
	signupGroup := echoInstance.Group("/signup")
	signupGroup.POST("", signupHandler.postSignup)
}

func enforceValidSession() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(session.CustomClaims)
		},
		ContextKey:    "Authorization",
		SigningMethod: "RS256",
		SigningKey:    session.SessionData().PublicKey,
	}
	return echojwt.WithConfig(config)
}
