package models

import "time"

// ArticleModel A model that describes an article
// @Description A model that describes an article
type ArticleModel struct {
	ArticleId    string    `json:"articleId" firestore:"articleId"`
	Title        string    `json:"title" firestore:"title"`
	Description  string    `json:"description" firestore:"description"`
	Slug         string    `json:"slug" firestore:"slug"`
	ContentURL   string    `json:"contentUrl" firestore:"contentUrl"`
	ThumbnailURL string    `json:"thumbnailUrl" firestore:"thumbnailUrl"`
	ThumbnailId  string    `json:"-" firestore:"thumbnailId"`
	ContentId    string    `json:"-" firestore:"contentId"`
	CreatedDate  time.Time `json:"created" firestore:"created"`
	UpdatedDate  time.Time `json:"updated" firestore:"updated"`
	DocumentRef  string    `json:"documentRef" firestore:"-"`
	Published    bool      `json:"-" firestore:"published"`
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
