// Package schema provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package schema

import (
	"github.com/labstack/echo/v4"
)

// CSRFModel defines model for CSRFModel.
type CSRFModel struct {
	Csrf string `json:"csrf"`
}

// DefaultResponseModel defines model for DefaultResponseModel.
type DefaultResponseModel struct {
	Message string `json:"message"`
}

// SignInModel defines model for SignInModel.
type SignInModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUpModel defines model for SignUpModel.
type SignUpModel struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserReadModel defines model for UserReadModel.
type UserReadModel struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
}

// SignInJSONRequestBody defines body for SignIn for application/json ContentType.
type SignInJSONRequestBody = SignInModel

// SignUpJSONRequestBody defines body for SignUp for application/json ContentType.
type SignUpJSONRequestBody = SignUpModel

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /api/auth/csrf)
	GetCsrfToken(ctx echo.Context) error

	// (POST /api/auth/signin)
	SignIn(ctx echo.Context) error

	// (GET /api/auth/signout)
	SignOut(ctx echo.Context) error

	// (POST /api/auth/signup)
	SignUp(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetCsrfToken converts echo context to params.
func (w *ServerInterfaceWrapper) GetCsrfToken(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetCsrfToken(ctx)
	return err
}

// SignIn converts echo context to params.
func (w *ServerInterfaceWrapper) SignIn(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SignIn(ctx)
	return err
}

// SignOut converts echo context to params.
func (w *ServerInterfaceWrapper) SignOut(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SignOut(ctx)
	return err
}

// SignUp converts echo context to params.
func (w *ServerInterfaceWrapper) SignUp(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SignUp(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/auth/csrf", wrapper.GetCsrfToken)
	router.POST(baseURL+"/api/auth/signin", wrapper.SignIn)
	router.GET(baseURL+"/api/auth/signout", wrapper.SignOut)
	router.POST(baseURL+"/api/auth/signup", wrapper.SignUp)

}
