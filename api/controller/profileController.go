package controller

import (
	"Clean_Architecture/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

func (profileController *ProfileController) Fetch(c *gin.Context) {
	userId := c.GetString("x-user-id")

	profile, err := profileController.ProfileUsecase.GetProfileByID(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
