package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"web-template/service"
	"web-template/util"
)

func ApiKeyAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		envApiKey := util.Config.ApiKey
		fullApiKey := c.GetHeader("Authorization")
		apiKey := strings.Split(fullApiKey, "Bearer ")

		if len(apiKey) == 2 && apiKey[1] == envApiKey {
			c.Status(200)
		} else {
			c.AbortWithStatusJSON(util.NewResponse(http.StatusUnauthorized, "invalid authentication", nil, nil))
			return
		}
	}
}

func BearerAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer, err := util.GetBearerToken(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(util.NewResponse(http.StatusUnauthorized, "invalid authentication", nil, nil))
			return
		}

		if service.ValidateBearerToken(bearer) {
			c.Status(200)
		} else {
			c.AbortWithStatusJSON(util.NewResponse(http.StatusUnauthorized, "invalid authentication", nil, nil))
			return
		}
	}
}
