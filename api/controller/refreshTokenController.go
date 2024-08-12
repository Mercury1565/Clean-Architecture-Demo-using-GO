package controller

import (
	"Clean_Architecture/bootstrap"
	"Clean_Architecture/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
	RefreshTokenUsecase domain.RefreshTokenUsecase
	Env                 *bootstrap.Env
}

func (refreshTokenController *RefreshTokenController) RefreshToken(c *gin.Context) {
	var request domain.RefreshTokenRequest

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := refreshTokenController.RefreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, refreshTokenController.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := refreshTokenController.RefreshTokenUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessTokenExp := refreshTokenController.Env.AccessTokenExpiryHour
	refreshTokenExp := refreshTokenController.Env.RefreshTokenExpiryHour
	accessTokenSecret := refreshTokenController.Env.AccessTokenSecret
	refreshTokenSecret := refreshTokenController.Env.RefreshTokenSecret

	accessToken, err := refreshTokenController.RefreshTokenUsecase.CreateAccessToken(&user, accessTokenSecret, accessTokenExp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := refreshTokenController.RefreshTokenUsecase.CreateRefreshToken(&user, refreshTokenSecret, refreshTokenExp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := domain.RefershTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
