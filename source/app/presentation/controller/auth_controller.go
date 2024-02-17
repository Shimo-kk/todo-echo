package controller

import (
	"net/http"
	"os"
	"time"
	"todo/app/core"
	"todo/app/presentation"
	"todo/app/service/schema"
	"todo/app/service/usecase"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type authController struct {
	usecase usecase.IAuthUsecase
}

// 認証コントローラーの作成
func NewAuthController(usecase usecase.IAuthUsecase) *authController {
	return &authController{usecase}
}

// CSRFトークンの取得
func (ac *authController) GetCsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)

	resonseBody := schema.CSRFModel{
		Csrf: token,
	}
	return c.JSON(http.StatusBadRequest, resonseBody)
}

// サインアップ
func (ac *authController) SignUp(c echo.Context) error {
	requestBody := schema.SignUpModel{}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, presentation.NewDefaultRespoce(err.Error()))
	}

	// ユースケース実行
	if err := ac.usecase.SignUp(requestBody); err != nil {
		dstErr := core.AsAppError(err)
		return c.JSON(presentation.ConvertErrorCode(dstErr.Code()), presentation.NewDefaultRespoce(dstErr.Error()))
	}

	return c.JSON(http.StatusOK, presentation.NewDefaultRespoce("サインアップに成功しました。"))
}

// サインイン
func (ac *authController) SignIn(c echo.Context) error {
	requestBody := schema.SignInModel{}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, presentation.NewDefaultRespoce(err.Error()))
	}

	// ユースケース実行
	responseBody, err := ac.usecase.SignIn(requestBody)
	if err != nil {
		dstErr := core.AsAppError(err)
		return c.JSON(presentation.ConvertErrorCode(dstErr.Code()), presentation.NewDefaultRespoce(dstErr.Error()))
	}

	// JWTトークンの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": responseBody.Id,
		"exp":     time.Now().Add(12 * time.Hour).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, presentation.NewDefaultRespoce(err.Error()))
	}

	// CookieにJWTトークンを設定
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenStr
	cookie.Expires = (time.Now().Add(24 * time.Hour))
	cookie.Path = "/"
	cookie.Domain = os.Getenv("MY_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, responseBody)
}

// サインアウト
func (ac *authController) SignOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("MY_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, presentation.NewDefaultRespoce("サインアウトしました。"))
}
