package middleware

import (
	"blog-backend/service"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiKeyAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		envApiKey := util.Config.ApiKey
		fullApiKey := c.GetHeader("X-Api-Key")

		if fullApiKey == envApiKey {
			c.Status(200)
		} else {
			response := util.NewResponse(http.StatusUnauthorized, "invalid authentication", "", nil)
			c.AbortWithStatusJSON(response.GetStatusCode(), response)
			return
		}
	}
}

func BearerAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer, err := util.GetBearerToken(c.Request)
		if err != nil {
			response := util.NewResponse(http.StatusUnauthorized, "invalid authentication", "", nil)
			c.AbortWithStatusJSON(response.GetStatusCode(), response)
			return
		}

		if service.ValidateBearerToken(bearer) {
			c.Status(200)
		} else {
			response := util.NewResponse(http.StatusUnauthorized, "invalid authentication", "", nil)
			c.AbortWithStatusJSON(response.GetStatusCode(), response)
			return
		}
	}
}
