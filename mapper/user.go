package mapper

import (
	"blog-backend/models"
	"github.com/bmwadforth-com/armor-go/src/util"
	"github.com/bmwadforth-com/armor-go/src/util/crypto"
	"github.com/google/uuid"
	"time"
)

func MapUserCreateRequest(request models.CreateUserRequest) models.UserModel {
	timeNow := time.Now()

	password, err := crypto.HashPassword([]byte(request.Password))
	if err != nil {
		util.LogError("failed to hash password: %v", err)
		return models.UserModel{}
	}

	userModel := models.UserModel{
		UserId:      uuid.New().String(),
		Username:    request.Username,
		Password:    string(password),
		CreatedDate: timeNow,
		UpdatedDate: timeNow,
		Active:      true,
	}

	return userModel
}
