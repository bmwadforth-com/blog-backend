package controllers

import (
	"blog-backend/grpc/protocol_buffers/gemini_service"
	"blog-backend/service"
	"blog-backend/util"
	"context"
	"fmt"
	armorUtil "github.com/bmwadforth-com/armor-go/src/util"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io"
	"net/http"
	"strings"
	"time"
)

func attachAPIKey(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "api-key", util.Config.ApiKey)
}

func attachTokenToContext(ctx context.Context, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
}

// QueryGemini example godoc
// @Summary Query gemini via gRPC
// @Schemes
// @Description Query gemini
// @Tags Query gemini
// @Accept json
// @Produce json
// @Success 200 {object}  util.ApiResponse[string]
// @Router /gemini [get]
func QueryGemini(c *gin.Context) {
	conn, err := grpc.NewClient(util.Config.GeminiService, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		armorUtil.LogFatal("%v", err)
	}
	defer conn.Close()

	client := gemini.NewGeminiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	ctxWithAPIKey := attachAPIKey(ctx)

	token, err := service.FetchIdentityToken(ctxWithAPIKey, fmt.Sprintf("https://%s", strings.Replace(util.Config.GeminiService, ":443", "", 1)))
	if err != nil {
		response := util.NewResponse(http.StatusInternalServerError, err.Error(), "", nil)
		c.JSON(response.GetStatusCode(), response)
		return
	}
	ctxWithIDToken := attachTokenToContext(ctxWithAPIKey, token)

	r, err := client.QueryGemini(ctxWithIDToken, &gemini.QueryRequest{Query: c.Query("query")})
	if err != nil {
		armorUtil.LogFatal("%v", err)
	}

	var messages []string
	for {
		message, err := r.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			response := util.NewResponse(http.StatusInternalServerError, err.Error(), "", nil)
			c.JSON(response.GetStatusCode(), response)
			return
		}

		messages = append(messages, message.Response)
	}

	c.Header("Content-Type", "text/markdown")
	c.String(http.StatusOK, strings.Join(messages, ""))
}
