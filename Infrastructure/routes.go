package router

import (
	controller "auth-server/Infrastructure/Controllers"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, ac controller.IAuthController) *echo.Echo {
	e := echo.New()

	e.GET("/login", uc.Login)
	e.GET("/logout", uc.Logout)
	e.GET("/signup", uc.SignUp)
	e.GET("/auth", ac.Auth)
	e.GET("/refresh", ac.Refresh)

	return e
}
