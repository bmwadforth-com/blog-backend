package util

import (
	"context"
	"encoding/json"
	"github.com/sethvargo/go-envconfig"
	"os"
	"path/filepath"
)

var Config configuration
var IsProduction bool

type configuration struct {
	ProjectId         string `json:"ProjectId" env:"WEB_TEMPLATE__PROJECTID"`
	ApiKey            string `json:"ApiKey" env:"WEB_TEMPLATE__APIKEY"`
	JwtSigningKey     string `json:"jwtSigningKey" env:"WEB_TEMPLATE__JWTSIGNINGKEY"`
	FireStoreDatabase string `json:"fireStoreDatabase" env:"WEB_TEMPLATE__FIRESTOREDATABASE"`
	Database          struct {
		Host     string `json:"host" env:"WEB_TEMPLATE__DATABASE_HOST"`
		Name     string `json:"name" env:"WEB_TEMPLATE__DATABASE_NAME"`
		Username string `json:"user" env:"WEB_TEMPLATE__DATABASE_USERNAME"`
		Password string `json:"pass" env:"WEB_TEMPLATE__DATABASE_PASSWORD"`
		Port     string `json:"port" env:"WEB_TEMPLATE__DATABASE_PORT"`
		SSL      bool   `json:"SSL" env:"WEB_TEMPLATE__DATABASE_SSL"`
	} `json:"database"`
}

func LoadConfiguration() {
	configFile, err := filepath.Abs("config.json")
	if err != nil {
		SLogger.Fatalf("an error has occurred: %v", err)
	}

	bytes, err := os.ReadFile(configFile)
	if err != nil {
		SLogger.Fatalf("an error has occurred: %v", err)
	}

	err = json.Unmarshal(bytes, &Config)
	if err != nil {
		SLogger.Fatalf("an error has occurred: %v", err)
	}
}

func LoadEnvironmentVariables() {
	ctx := context.Background()

	if err := envconfig.Process(ctx, &Config); err != nil {
		SLogger.Fatalf("an error has occurred: %v", err)
	}
}
