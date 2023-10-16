package models

import "time"

// ArticleModel A model that describes an article
// @Description A model that describes an article
type ArticleModel struct {
	ArticleId   string    `json:"articleId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ThumbnailId string    `json:"thumbnailId"`
	ContentId   string    `json:"contentId"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	Published   bool
}

// CreateArticleRequest New article
// @Description New article request
type CreateArticleRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ThumbnailId string `json:"thumbnailId"`
	ContentId   string `json:"contentId"`
}
