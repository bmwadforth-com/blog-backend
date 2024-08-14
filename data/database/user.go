package database

import (
	"blog-backend/data/mapper"
	"blog-backend/data/models"
	"blog-backend/util"
	"context"
	"errors"
)

func CreateUser(request models.CreateUserRequest, ctx context.Context) util.DataResponse[string] {
	dataResponse := util.NewDataResponse("successfully created user", "")

	user := mapper.MapUserCreateRequest(request)
	_, _, err := DbConnection.Collection("users").Add(ctx, user)
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	dataResponse.SetData(user.UserId)

	return dataResponse
}

func GetUserByUsername(username string, ctx context.Context) util.DataResponse[models.UserModel] {
	var user models.UserModel
	dataResponse := util.NewDataResponse("success", user)

	docs, err := DbConnection.Collection("users").Where("Username", "==", username).Documents(ctx).GetAll()
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	if len(docs) == 0 {
		dataResponse.SetError(errors.New("no users found"), util.DbresultNotFound)
		return dataResponse
	}

	if len(docs) > 1 {
		dataResponse.SetError(errors.New("error multiple users found"), util.DbresultError)
		return dataResponse
	}

	err = docs[0].DataTo(&user)
	user.DocumentRef = docs[0].Ref.ID
	if err != nil {
		dataResponse.SetError(errors.New("error unable to deserialize record"), util.DbresultError)
		return dataResponse
	}
	dataResponse.SetData(user)

	return dataResponse
}
