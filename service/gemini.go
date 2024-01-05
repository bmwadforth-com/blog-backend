package service

import (
	GeminiService "blog-backend/protocol_buffers/gemini_service"
	"blog-backend/util"
	"context"
	"errors"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Server struct {
	GeminiService.UnimplementedGeminiServer
}

func (s *Server) QueryGemini(in *GeminiService.QueryRequest, srv GeminiService.Gemini_QueryGeminiServer) error {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(util.Config.GeminiApiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-and-image input (multimodal), use the gemini-pro-vision model. gemini-pro for text based only
	model := client.GenerativeModel("gemini-pro")

	prompt := genai.Text(in.Query)
	iter := model.GenerateContentStream(ctx, prompt)

	for {
		resp, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(resp.Candidates) > 1 {
			// TODO: More than one candidate response to chose from. Do something
		} else {
			var response string
			for _, part := range resp.Candidates[0].Content.Parts {
				switch part.(type) {
				case genai.Text:
					response = string(part.(genai.Text))
				case genai.Blob:
					util.SLogger.Errorf("Received invalid type of data from Gemini response")
					continue
				default:
					util.SLogger.Errorf("Received invalid type of data from Gemini response")
					continue
				}

				if err := srv.Send(&GeminiService.QueryResponse{Response: response}); err != nil {
					util.SLogger.Errorf("Error occurred: %v", err)
					err := status.Error(codes.Internal, err.Error())
					return err
				}
			}
		}
	}

	return nil
}
