package middlewares

import (
	"net/http"
	"video-storage/configs"
	"video-storage/types"
	"video-storage/utils"

	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	t, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, types.CustomResponse{Ok: false, Message: err.Error()})
		c.Abort()
		return
	}
	authorized, err := utils.IsAuthorized(t, configs.ENV.JwtSecret)
	if authorized {
		payload, err := utils.ExtractPayloadFromToken(t, configs.ENV.JwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, types.CustomResponse{Ok: false, Message: err.Error()})
			c.Abort()
			return
		}
		c.Set("user_id", payload)
		c.Next()
		return
	}
	c.JSON(http.StatusUnauthorized, types.CustomResponse{Ok: false, Message: err.Error()})
	c.Abort()
}
