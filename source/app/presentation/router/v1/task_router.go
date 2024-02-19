package v1

import (
	"todo/app/infrastructure/repository"
	"todo/app/presentation/controller"
	"todo/app/service/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func IncludeTaskRouter(db *gorm.DB, g *echo.Group) {
	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	g.POST("/task", taskController.Create)
	g.PUT("/task", taskController.Update)
	g.GET("/task/done/:id", taskController.Done)
	g.GET("/task/:id", taskController.Get)
	g.DELETE("/task/:id", taskController.Delete)
	g.GET("/tasks", taskController.GetAll)
}
