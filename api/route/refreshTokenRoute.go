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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, database mongo.Database, group *gin.RouterGroup) {
	userRepo := repository.NewUserRepo(database, domain.CollectionUser)
	refershTokenController := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(userRepo, timeout),
		Env:                 bootstrap.NewEnv(),
	}

	group.POST("/refresh", refershTokenController.RefreshToken)
}
