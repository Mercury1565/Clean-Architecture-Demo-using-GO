package middelware

import (
	"Clean_Architecture/domain"
	"Clean_Architecture/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		splitted := strings.Split(authHeader, " ")

		if len(splitted) == 2 {
			authToken := splitted[1]
			authorized, err := utils.IsAuthorized(authToken, secret)

			if authorized {
				userId, err := utils.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-id", userId)
				c.Next()
				return
			}

			// not authorized
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		// invalid authorization header
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "unauthorized user"})
		c.Abort()
	}
}
