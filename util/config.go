package util

import (
	armor "github.com/bmwadforth-com/armor-go/src"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

var Config configuration
var IsProduction bool

type configuration struct {
	ProjectId          string `json:"ProjectId" env:"WEB_TEMPLATE__PROJECTID"`
	ApiKey             string `json:"ApiKey" env:"WEB_TEMPLATE__APIKEY"`
	GeminiService      string `json:"GeminiService" env:"WEB_TEMPLATE__GEMINISERVICE"`
	JwtSigningKey      string `json:"jwtSigningKey" env:"WEB_TEMPLATE__JWTSIGNINGKEY"`
	FireStoreDatabase  string `json:"fireStoreDatabase" env:"WEB_TEMPLATE__FIRESTOREDATABASE"`
	CloudStorageBucket string `json:"cloudStorageBucket" env:"WEB_TEMPLATE__CLOUDSTORAGEBUCKET"`
	ContentURL         string `json:"contentURL" env:"WEB_TEMPLATE__CONTENTURL"`
}

func (c configuration) Validate() error {
	return nil
}

func SetupArmor() error {
	IsProduction = os.Getenv("APP_ENV") == "PRODUCTION"

	localConfigFile, err := filepath.Abs("config.local.json")
	if err != nil {
		return err
	}
	_, err = os.Stat(localConfigFile)
	if err != nil {
		return err
	}

	return armor.InitArmor(IsProduction, zapcore.InfoLevel, &Config, localConfigFile)
}
