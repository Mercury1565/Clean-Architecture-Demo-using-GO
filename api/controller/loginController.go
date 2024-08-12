package controller

import (
	"Clean_Architecture/bootstrap"
	"Clean_Architecture/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase domain.LoginUseCase
	Env          *bootstrap.Env
}

func (loginController *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := loginController.LoginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessTokenExp := loginController.Env.AccessTokenExpiryHour
	refreshTokenExp := loginController.Env.RefreshTokenExpiryHour
	accessTokenSecret := loginController.Env.AccessTokenSecret
	refreshTokenSecret := loginController.Env.RefreshTokenSecret

	accessToken, err := loginController.LoginUsecase.CreateAccessToken(&user, accessTokenSecret, accessTokenExp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := loginController.LoginUsecase.CreateRefreshToken(&user, refreshTokenSecret, refreshTokenExp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
