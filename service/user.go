package service

import (
	"blog-backend/database"
	"blog-backend/models"
	"blog-backend/util"
	"errors"
	"time"
)

var Sessions = map[string]models.UserSessionModel{}

func comparePassword(password string, hashedPassword string) bool {
	passwordMatch := util.PasswordHashMatch([]byte(hashedPassword), []byte(password))

	if passwordMatch {
		return true
	}

	return false
}

func LoginUser(request models.LoginUserRequest) (bool, string, error) {
	dataResponse := database.GetUserByUsername(request.Username)

	if dataResponse.GetError() != nil {
		return false, "", dataResponse.GetError()
	}

	validPassword := comparePassword(request.Password, dataResponse.Data.Password)
	if validPassword {
		token := NewBearerToken(dataResponse.Data.Username)

		session := models.UserSessionModel{
			UserId:   dataResponse.Data.UserId,
			Username: dataResponse.Data.Username,
			LoggedIn: time.Now(),
			Token:    string(token),
			Active:   true,
		}

		Sessions[dataResponse.Data.Username] = session

		return true, string(token), nil
	}

	return false, "", errors.New("challenge failed")
}
