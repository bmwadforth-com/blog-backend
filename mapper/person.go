package mapper

import (
	"github.com/google/uuid"
	"blog-backend/models"
)

func MapPersonCreateRequest(request models.PersonCreateRequest) models.PersonModel {
	return models.PersonModel{
		Identifier:  uuid.New().String(),
		FirstName:   request.FirstName,
		MiddleName:  request.MiddleName,
		LastName:    request.LastName,
		DateOfBirth: request.DateOfBirth,
	}
}
