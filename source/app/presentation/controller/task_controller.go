package controller

import (
	"net/http"
	"strconv"
	"todo/app/core"
	"todo/app/presentation"
	"todo/app/service/schema"
	"todo/app/service/usecase"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type taskController struct {
	usecase usecase.ITaskUsecase
}

// タスクコントローラーの作成
func NewTaskController(usecase usecase.ITaskUsecase) *taskController {
	return &taskController{usecase: usecase}
}

// タスクを作成
func (tc *taskController) Create(c echo.Context) error {
	requestBody := schema.TaskCreateModel{}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, presentation.NewDefaultRespoce(err.Error()))
	}

	// JWTトークンからユーザーIDを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// ユースケース実行
	responseBody, err := tc.usecase.CreateTask(int(userId.(float64)), requestBody)
	if err != nil {
		dstErr := core.AsAppError(err)
		return c.JSON(presentation.ConvertErrorCode(dstErr.Code()), presentation.NewDefaultRespoce(dstErr.Error()))
	}

	return c.JSON(http.StatusOK, responseBody)
}

// タスクを取得
func (tc *taskController) Get(c echo.Context) error {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	// ユースケース実行
	responseBody, err := tc.usecase.GetTask(int(id))
	if err != nil {
		dstErr := core.AsAppError(err)
		return c.JSON(presentation.ConvertErrorCode(dstErr.Code()), presentation.NewDefaultRespoce(dstErr.Error()))
	}

	return c.JSON(http.StatusOK, responseBody)
}

// タスクを全件取得
func (tc *taskController) GetAll(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// ユースケース実行
	responseBody, err := tc.usecase.GetAllTask(int(userId.(float64)))
	if err != nil {
		dstErr := core.AsAppError(err)
		return c.JSON(presentation.ConvertErrorCode(dstErr.Code()), presentation.NewDefaultRespoce(dstErr.Error()))
	}

	return c.JSON(http.StatusOK, responseBody)
}

// タスクを更新
func (tc *taskController) Update(c echo.Context) error {
	requestBody := schema.TaskUpdateModel{}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, presentation.NewDefaultRespoce(err.Error()))
	}

	// ユースケース実行
	responseBody, err := tc.usecase.UpdateTask(requestBody)
	if err != nil {
		dstErr := core.AsAppError(err)
		return c.JSON(presentation.ConvertErrorCode(dstErr.Code()), presentation.NewDefaultRespoce(dstErr.Error()))
	}

	return c.JSON(http.StatusOK, responseBody)
}

// タスクを完了
func (tc *taskController) Done(c echo.Context) error {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	// ユースケース実行
	responseBody, err := tc.usecase.DoneTask(int(id))
	if err != nil {
		dstErr := core.AsAppError(err)
		return c.JSON(presentation.ConvertErrorCode(dstErr.Code()), presentation.NewDefaultRespoce(dstErr.Error()))
	}

	return c.JSON(http.StatusOK, responseBody)
}

// タスクを削除
func (tc *taskController) Delete(c echo.Context) error {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	// ユースケース実行
	err := tc.usecase.DeleteTask(int(id))
	if err != nil {
		dstErr := core.AsAppError(err)
		return c.JSON(presentation.ConvertErrorCode(dstErr.Code()), presentation.NewDefaultRespoce(dstErr.Error()))
	}

	return c.JSON(http.StatusOK, presentation.NewDefaultRespoce("タスクを削除しました。"))
}
