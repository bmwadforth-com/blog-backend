package mapper

import (
	"blog-backend/models"
	"blog-backend/util"
	"github.com/google/uuid"
	"time"
)

func MapUserCreateRequest(request models.CreateUserRequest) models.UserModel {
	timeNow := time.Now()

	userModel := models.UserModel{
		UserId:      uuid.New().String(),
		Username:    request.Username,
		Password:    string(util.HashPassword([]byte(request.Password))),
		CreatedDate: timeNow,
		UpdatedDate: timeNow,
		Active:      true,
	}

	return userModel
}
