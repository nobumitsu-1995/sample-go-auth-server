package controller

import "github.com/labstack/echo/v4"

type IUserController interface {
	Login(c echo.Context) error
	Logout(c echo.Context) error
	SignUp(c echo.Context) error
}
