package models

import "time"

// ArticleModel A model that describes an article
// @Description A model that describes an article
type ArticleModel struct {
	ArticleId    string    `json:"articleId"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Slug         string    `json:"slug"`
	ContentURL   string    `json:"contentUrl"`
	ThumbnailURL string    `json:"thumbnailUrl"`
	ThumbnailId  string    `json:"-"`
	ContentId    string    `json:"-"`
	CreatedDate  time.Time `json:"created"`
	UpdatedDate  time.Time `json:"updated"`
	DocumentRef  string    `json:"documentRef"`
	Published    bool      `json:"-"`
}

// CreateArticleRequest New article
// @Description New article request
type CreateArticleRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// CreateArticleContentResponse Article content/thumbnail creation response
// @Article content/thumbnail creation response
type CreateArticleContentResponse struct {
	ThumbnailId string `json:"thumbnailId"`
	ContentId   string `json:"contentId"`
}
