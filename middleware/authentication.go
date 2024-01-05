package middleware

import (
	"blog-backend/service"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
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

func BearerAuthenticationInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return status.Errorf(codes.InvalidArgument, "missing metadata")
	}

	// Check for the authorization metadata
	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	// Expecting the value format to be "Bearer <token>"
	token := values[0]

	// Expecting the value format to be "<token>"
	tokenString := strings.Replace(token, "Bearer ", "", 1)

	if !service.ValidateBearerToken(tokenString) {
		return status.Errorf(codes.Unauthenticated, "invalid JWT token")
	}

	// Token is valid, proceed with the handler
	return handler(srv, ss)
}
