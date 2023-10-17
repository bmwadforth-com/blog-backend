package controllers

import (
	"blog-backend/database"
	"blog-backend/models"
	"blog-backend/storage"
	"blog-backend/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateArticle example godoc
// @Summary Create article
// @Schemes
// @Description Create article
// @Tags Create article
// @Param Create article body models.CreateArticleRequest true "Create article model object"
// @Accept json
// @Produce json
// @Success 200 {object}  util.ApiResponse
// @Router /article [post]
func CreateArticle(c *gin.Context) {
	var request models.CreateArticleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := util.NewResponse(http.StatusBadRequest, "Failed to parse JSON", "", err)
		c.JSON(response.GetStatusCode(), response)
		return
	}

	r := database.CreateArticle(request)
	if r.GetError() != nil {
		response := util.NewResponse(http.StatusInternalServerError, "An error has occurred", "", r.GetError())
		c.JSON(response.GetStatusCode(), response)
		return
	}

	response := util.NewResponse(http.StatusOK, r.Message, r.Data, nil)
	c.JSON(response.GetStatusCode(), response)
}

// GetArticles example godoc
// @Summary Get articles
// @Schemes
// @Description Get articles
// @Tags Get articles
// @Accept json
// @Produce json
// @Success 200 {object}  util.ApiResponse
// @Router /articles [get]
func GetArticles(c *gin.Context) {
	r := database.GetArticles()
	if r.GetError() != nil {
		response := util.NewResponse(http.StatusInternalServerError, "An error has occurred", "", r.GetError())
		c.JSON(response.GetStatusCode(), response)
		return
	}

	response := util.NewResponse(http.StatusOK, r.Message, r.Data, nil)
	c.JSON(response.GetStatusCode(), response)
}

func UploadArticle(c *gin.Context) {
	articleId := c.Param("articleId")
	contentFileHeader, _ := c.FormFile("content")
	thumbnailFileHeader, _ := c.FormFile("thumbnail")

	var contentFileMultiPart *storage.MultipartFile
	var thumbnailFileMultiPart *storage.MultipartFile

	if contentFileHeader == nil && thumbnailFileHeader == nil {
		response := util.NewResponse(http.StatusBadRequest, "Bad request", "", errors.New("supply either content or thumbnail files in multipart form body"))
		c.JSON(response.GetStatusCode(), response)
		return
	}

	if contentFileHeader != nil {
		contentFile, err := contentFileHeader.Open()
		if err != nil {
			response := util.NewResponse(http.StatusInternalServerError, "An error has occurred", "", err)
			c.JSON(response.GetStatusCode(), response)
			return
		}

		contentFileMultiPart = &storage.MultipartFile{
			File:     contentFile,
			FileSize: contentFileHeader.Size,
		}
	}

	if thumbnailFileHeader != nil {
		thumbnailFile, err := thumbnailFileHeader.Open()
		if err != nil {
			response := util.NewResponse(http.StatusInternalServerError, "An error has occurred", "", err)
			c.JSON(response.GetStatusCode(), response)
			return
		}

		thumbnailFileMultiPart = &storage.MultipartFile{
			File:     thumbnailFile,
			FileSize: thumbnailFileHeader.Size,
		}
	}

	contentId, thumbnailId, err := storage.UploadArticleContent(articleId, contentFileMultiPart, thumbnailFileMultiPart)
	if err != nil {
		response := util.NewResponse(http.StatusInternalServerError, "An error has occurred", "", err)
		c.JSON(response.GetStatusCode(), response)
		return
	}

	apiResponse := models.CreateArticleContentResponse{
		ContentId:   contentId,
		ThumbnailId: thumbnailId,
	}

	response := util.NewResponse(http.StatusOK, "successfully uploaded content", apiResponse, nil)
	c.JSON(response.GetStatusCode(), response)
}
