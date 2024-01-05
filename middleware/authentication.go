package middleware

import (
	"blog-backend/service"
	"blog-backend/util"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

func BearerAuthenticationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "retrieving metadata failed")
	}

	tokenSlice, ok := md["authorization"]
	if !ok || len(tokenSlice) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "JWT token is not provided")
	}

	tokenString := tokenSlice[0]
	if !service.ValidateBearerToken(tokenString) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid JWT token")
	}

	return handler(ctx, req)
}
