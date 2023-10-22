package database

import (
	"blog-backend/mapper"
	"blog-backend/models"
	"blog-backend/util"
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"time"
)

func GetArticle(articleId string) util.DataResponse[models.ArticleModel] {
	var article models.ArticleModel
	dataResponse := util.NewDataResponse("success", article)
	ctx := context.Background()
	client, _ := createClient(ctx)
	defer client.Close()

	docs, err := client.Collection("articles").Where("ArticleId", "==", articleId).Where("Published", "==", true).Documents(ctx).GetAll()
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	if len(docs) == 0 {
		dataResponse.SetError(errors.New("no articles found"), util.DbresultNotFound)
		return dataResponse
	}

	if len(docs) > 1 {
		dataResponse.SetError(errors.New("error multiple articles found"), util.DbresultError)
		return dataResponse
	}

	err = docs[0].DataTo(&article)
	article.DocumentRef = docs[0].Ref.ID
	if article.ContentId != "" {
		article.ContentURL = fmt.Sprintf("%s/%s/%s", util.Config.ContentURL, article.Slug, article.ContentId)
	}

	if article.ThumbnailId != "" {
		article.ThumbnailURL = fmt.Sprintf("%s/%s/%s", util.Config.ContentURL, article.Slug, article.ThumbnailId)
	}

	if err != nil {
		dataResponse.SetError(errors.New("error unable to deserialize record"), util.DbresultError)
		return dataResponse
	}
	dataResponse.SetData(article)

	return dataResponse
}

func GetArticleBySlug(slug string) util.DataResponse[models.ArticleModel] {
	var article models.ArticleModel
	dataResponse := util.NewDataResponse("success", article)
	ctx := context.Background()
	client, _ := createClient(ctx)
	defer client.Close()

	docs, err := client.Collection("articles").Where("Slug", "==", slug).Where("Published", "==", true).Documents(ctx).GetAll()
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	if len(docs) == 0 {
		dataResponse.SetError(errors.New("no articles found"), util.DbresultNotFound)
		return dataResponse
	}

	if len(docs) > 1 {
		dataResponse.SetError(errors.New("error multiple articles found"), util.DbresultError)
		return dataResponse
	}

	err = docs[0].DataTo(&article)
	article.DocumentRef = docs[0].Ref.ID
	if article.ContentId != "" {
		article.ContentURL = fmt.Sprintf("%s/%s/%s", util.Config.ContentURL, article.Slug, article.ContentId)
	}

	if article.ThumbnailId != "" {
		article.ThumbnailURL = fmt.Sprintf("%s/%s/%s", util.Config.ContentURL, article.Slug, article.ThumbnailId)
	}

	if err != nil {
		dataResponse.SetError(errors.New("error unable to deserialize record"), util.DbresultError)
		return dataResponse
	}
	dataResponse.SetData(article)

	return dataResponse
}

func GetArticles() util.DataResponse[[]models.ArticleModel] {
	var articles []models.ArticleModel
	dataResponse := util.NewDataResponse("successfully read articles", articles)
	ctx := context.Background()
	client, _ := createClient(ctx)
	defer client.Close()

	docs, err := client.Collection("articles").Where("Published", "==", true).Documents(ctx).GetAll()
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	for _, doc := range docs {
		article := models.ArticleModel{}
		doc.DataTo(&article)
		article.DocumentRef = doc.Ref.ID
		if article.ContentId != "" {
			article.ContentURL = fmt.Sprintf("%s/%s/%s", util.Config.ContentURL, article.Slug, article.ContentId)
		}

		if article.ThumbnailId != "" {
			article.ThumbnailURL = fmt.Sprintf("%s/%s/%s", util.Config.ContentURL, article.Slug, article.ThumbnailId)
		}

		articles = append(articles, article)
	}

	dataResponse.SetData(articles)

	return dataResponse
}

func CreateArticle(request models.CreateArticleRequest) util.DataResponse[string] {
	dataResponse := util.NewDataResponse("successfully created article", "")
	ctx := context.Background()
	client, _ := createClient(ctx)
	defer client.Close()

	article := mapper.MapArticleCreatRequest(request)
	_, _, err := client.Collection("articles").Add(ctx, article)
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	dataResponse.SetData(article.ArticleId)

	return dataResponse
}

func UpdateArticle(article models.ArticleModel) util.DataResponse[string] {
	dataResponse := util.NewDataResponse("successfully updated article", "")
	ctx := context.Background()
	client, _ := createClient(ctx)
	defer client.Close()

	_, err := client.Collection("articles").Doc(article.DocumentRef).Set(ctx, map[string]interface{}{
		"Title":       article.Title,
		"Description": article.Description,
		"Slug":        article.Slug,
		"Published":   article.Published,
		"ContentId":   article.ContentId,
		"ThumbnailId": article.ThumbnailId,
		"UpdatedDate": time.Now(),
	}, firestore.MergeAll)
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	dataResponse.SetData(article.ArticleId)

	return dataResponse
}
