package storage

import (
	"blog-backend/database"
	"blog-backend/util"
	"context"
	"fmt"
	"github.com/google/uuid"
	"mime/multipart"
)

type MultipartFile struct {
	File     multipart.File
	FileSize int64
}

func UploadArticleContent(articleId string, content *MultipartFile, thumbnail *MultipartFile, ctx context.Context) (string, string, error) {
	dataResponse := database.GetArticle(articleId, ctx)
	if dataResponse.GetError() != nil {
		return "", "", dataResponse.GetError()
	}
	article := dataResponse.Data

	if article.ContentId == "" || article.ThumbnailId == "" {
		if article.ContentId == "" && content != nil {
			article.ContentId = uuid.New().String()
			article.ContentURL = fmt.Sprintf("%s/%s/%s", util.Config.ContentURL, article.Slug, article.ContentId)
		}

		if article.ThumbnailId == "" && thumbnail != nil {
			article.ThumbnailId = uuid.New().String()
			article.ThumbnailURL = fmt.Sprintf("%s/%s/%s", util.Config.ContentURL, article.Slug, article.ThumbnailId)
		}

		database.UpdateArticle(article, ctx)
	}

	// Upload article content
	if content != nil {
		contentBytes := make([]byte, content.FileSize)
		_, err := content.File.Read(contentBytes)
		if err != nil {
			util.SLogger.Errorf("failed to upload article content: %v", err)
			return "", "", err
		}

		if util.IsProduction {
			err = streamFileUpload(fmt.Sprintf("%s/%s", article.Slug, article.ContentId), contentBytes)
		} else {
			err = streamFileUpload(fmt.Sprintf("development/%s/%s", article.Slug, article.ContentId), contentBytes)
		}
		if err != nil {
			return "", "", err
		}
	}

	// Upload article thumbnail
	if thumbnail != nil {
		thumbnailBytes := make([]byte, thumbnail.FileSize)
		_, err := thumbnail.File.Read(thumbnailBytes)
		if err != nil {
			util.SLogger.Errorf("failed to upload article thumbnail: %v", err)
			return "", "", err
		}

		if util.IsProduction {
			err = streamFileUpload(fmt.Sprintf("%s/%s", article.Slug, article.ThumbnailId), thumbnailBytes)
		} else {
			err = streamFileUpload(fmt.Sprintf("development/%s/%s", article.Slug, article.ThumbnailId), thumbnailBytes)
		}
		if err != nil {
			return "", "", err
		}
	}

	return article.ContentId, article.ThumbnailId, nil
}
