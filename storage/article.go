package storage

import (
	"blog-backend/database"
	"blog-backend/util"
	"github.com/google/uuid"
	"mime/multipart"
)

func UploadArticleContent(articleId string, file multipart.File, fileSize int64) (string, error) {
	dataResponse := database.GetArticle(articleId)
	article := dataResponse.Data
	contentId := article.ContentId
	if contentId == "" {
		contentId = uuid.New().String()
		article.ContentId = contentId
	}

	fileBytes := make([]byte, 0, fileSize)
	_, err := file.Read(fileBytes)
	if err != nil {
		util.SLogger.Errorf("failed to upload article content: %v", err)
		return "", err
	}

	err = streamFileUpload(contentId, fileBytes)
	if err != nil {
		return "", err
	}

	// TODO: write contentId back into article
	// database.UpdateArticle(articleId, article)

	return contentId, nil
}
