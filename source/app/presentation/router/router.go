package router

import (
	"net/http"
	"os"
	"todo/app/infrastructure/database/postgres"
	"todo/app/infrastructure/repository"
	"todo/app/presentation/controller"
	"todo/app/service/usecase"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func IncludeRouter(e *echo.Echo) {
	db := postgres.NewDB()
	userRepository := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepository)
	authController := controller.NewAuthController(authUsecase)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	// apiグループ
	api := e.Group("/api")

	// CORSミドルウェアの設定
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken,
		},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowCredentials: true,
	}))

	// CSRFミドルウェアの設定
	api.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("MY_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		CookieMaxAge:   60,
	}))

	// 認証関連
	api.GET("/auth/csrf", authController.GetCsrfToken)
	api.POST("/auth/signup", authController.SignUp)
	api.POST("/auth/signin", authController.SignIn)
	api.GET("/auth/signout", authController.SignOut)

	// v1グループ
	v1 := api.Group("/v1")

	// JWTミドルウェアの設定
	v1.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("JWT_KEY")),
		TokenLookup: "cookie:token",
	}))
}
