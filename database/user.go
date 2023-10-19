package database

import (
	"blog-backend/mapper"
	"blog-backend/models"
	"blog-backend/util"
	"context"
	"errors"
)

func CreateUser(request models.CreateUserRequest) util.DataResponse[string] {
	dataResponse := util.NewDataResponse("successfully created user", "")
	ctx := context.Background()
	client, _ := createClient(ctx)
	defer client.Close()

	user := mapper.MapUserCreateRequest(request)
	_, _, err := client.Collection("users").Add(ctx, user)
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	dataResponse.SetData(user.UserId)

	return dataResponse
}

func GetUserByUsername(username string) util.DataResponse[models.UserModel] {
	var user models.UserModel
	dataResponse := util.NewDataResponse("success", user)
	ctx := context.Background()
	client, _ := createClient(ctx)
	defer client.Close()

	docs, err := client.Collection("users").Where("Username", "==", username).Documents(ctx).GetAll()
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
