package controllers

import (
	pb "blog-backend/protocol_buffers/gemini_service"
	"blog-backend/util"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func attachAPIKey(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "api-key", util.Config.ApiKey)
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
	conn, err := grpc.Dial(util.Config.GeminiService, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGeminiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	ctxWithAPIKey := attachAPIKey(ctx)

	r, err := client.QueryGemini(ctxWithAPIKey, &pb.QueryRequest{Query: c.Query("query")})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
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
