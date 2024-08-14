package mapper

import (
	"blog-backend/data/models"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"time"
)

func MapArticleCreatRequest(request models.CreateArticleRequest) models.ArticleModel {
	timeNow := time.Now()
	sluggedTitle := slug.Make(request.Title)

	articleModel := models.ArticleModel{
		ArticleId:   uuid.New().String(),
		Title:       request.Title,
		Description: request.Description,
		Slug:        sluggedTitle,
		ThumbnailId: "",
		ContentId:   "",
		CreatedDate: timeNow,
		UpdatedDate: timeNow,
		Published:   false,
	}

	return articleModel
}
