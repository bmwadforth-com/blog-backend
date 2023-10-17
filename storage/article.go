package storage

import (
	"blog-backend/database"
	"blog-backend/util"
	"github.com/google/uuid"
	"mime/multipart"
)

type MultipartFile struct {
	File     multipart.File
	FileSize int64
}

func UploadArticleContent(articleId string, content *MultipartFile, thumbnail *MultipartFile) (string, string, error) {
	dataResponse := database.GetArticle(articleId)
	article := dataResponse.Data

	if article.ContentId == "" || article.ThumbnailId == "" {
		if article.ContentId == "" && content != nil {
			article.ContentId = uuid.New().String()
		}

		if article.ThumbnailId == "" && thumbnail != nil {
			article.ThumbnailId = uuid.New().String()
		}

		database.UpdateArticle(article)
	}

	// Upload article content
	if content != nil {
		contentBytes := make([]byte, content.FileSize)
		_, err := content.File.Read(contentBytes)
		if err != nil {
			util.SLogger.Errorf("failed to upload article content: %v", err)
			return "", "", err
		}
		err = streamFileUpload(article.ContentId, contentBytes)
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
		err = streamFileUpload(article.ThumbnailId, thumbnailBytes)
		if err != nil {
			return "", "", err
		}
	}

	return article.ContentId, article.ThumbnailId, nil
}
