package mapper

import (
	"blog-backend/models"
	"github.com/google/uuid"
	"time"
)

func MapArticleCreatRequest(request models.CreateArticleRequest) models.ArticleModel {
	timeNow := time.Now()

	articleModel := models.ArticleModel{
		ArticleId:   uuid.New().String(),
		Title:       request.Title,
		Description: request.Description,
		Slug:        request.Slug,
		ThumbnailId: "",
		ContentId:   "",
		CreatedDate: timeNow,
		UpdatedDate: timeNow,
		Published:   false,
	}

	if request.ThumbnailId != "" {
		articleModel.ThumbnailId = request.ThumbnailId
	}

	if request.ContentId != "" {
		articleModel.ContentId = request.ContentId
	}

	return articleModel
}
