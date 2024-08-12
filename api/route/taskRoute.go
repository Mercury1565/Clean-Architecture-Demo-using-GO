package route

import (
	"Clean_Architecture/api/controller"
	"Clean_Architecture/bootstrap"
	"Clean_Architecture/domain"
	"Clean_Architecture/repository"
	"Clean_Architecture/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, database mongo.Database, group *gin.RouterGroup) {
	taskRepo := repository.NewTaskRepo(database, domain.CollectionTask)
	taskController := &controller.TaskController{
		TaskUseCase: usecase.NewTaskUsecase(taskRepo, timeout),
	}

	group.POST("/task", taskController.Create)
	group.GET("/task", taskController.Fetch)
}
