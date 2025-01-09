package controller

import (
	"auth-server/Usecase/Services/Sessions"
	"auth-server/Usecase/Services/Users"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IAuthController interface {
	Auth(c echo.Context) error
	Refresh(c echo.Context) error
}

type AuthController struct {
	userService    Users.IUserService
	sessionService Sessions.ISessionService
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func NewAuthController(userService Users.IUserService, sessionService Sessions.ISessionService) IAuthController {
	return &AuthController{userService, sessionService}
}

func (ac *AuthController) Auth(c echo.Context) error {
	userId := c.Param("userId")

	err := ac.sessionService.ValidateSession(userId, c.Request().Header.Get("Authorization"))
	if err != nil {
		return c.String(http.StatusUnauthorized, "Invalid session")
	}

	return c.String(http.StatusOK, "Valid session")
}

func (ac *AuthController) Refresh(c echo.Context) error {
	userId := c.Param("userId")
	var request RefreshTokenRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	refreshToken := request.RefreshToken

	newAccessToken, newRefreshToken, err := ac.sessionService.RefreshSession(userId, refreshToken)
	if err != nil {
		return c.String(http.StatusUnauthorized, "Invalid session")
	}

	ac.setCookie(c, newAccessToken.Token)

	return c.JSON(http.StatusOK, map[string]string{"refreshToken": newRefreshToken.Token})
}

func (ac *AuthController) setCookie(c echo.Context, sessionId string) {
	cookie := new(http.Cookie)
	cookie.Name = "sessionId"
	cookie.Value = sessionId
	cookie.Expires = time.Now().Add(time.Minute * 15)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
}
