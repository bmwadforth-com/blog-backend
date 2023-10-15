package database

import (
	"context"
	"web-template/mapper"
	"web-template/models"
	"web-template/util"
)

func GetPeople() util.DataResponse[[]models.PersonModel] {
	var people []models.PersonModel
	var person models.PersonModel
	dataResponse := util.NewDataResponse("successfully read people records", people)
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	docs, _ := client.Collection("person").Documents(ctx).GetAll()

	for _, doc := range docs {
		doc.DataTo(&person)
		people = append(people, person)
	}

	dataResponse.Data = people

	return dataResponse
}

func CreatePerson(request models.PersonCreateRequest) util.DataResponse[string] {
	dataResponse := util.NewDataResponse("successfully wrote person record", "")
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	_, _, err := client.Collection("person").Add(ctx, mapper.MapPersonCreateRequest(request))
	if err != nil {
		util.SLogger.Errorf("an error occurred adding a person: %v", err)
		dataResponse.SetError(err)
	}

	return dataResponse
}
