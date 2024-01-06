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
	ProjectId          string `json:"ProjectId" env:"WEB_TEMPLATE__PROJECTID"`
	ApiKey             string `json:"ApiKey" env:"WEB_TEMPLATE__APIKEY"`
	GeminiService      string `json:"GeminiService" env:"WEB_TEMPLATE__GEMINISERVICE"`
	JwtSigningKey      string `json:"jwtSigningKey" env:"WEB_TEMPLATE__JWTSIGNINGKEY"`
	FireStoreDatabase  string `json:"fireStoreDatabase" env:"WEB_TEMPLATE__FIRESTOREDATABASE"`
	CloudStorageBucket string `json:"cloudStorageBucket" env:"WEB_TEMPLATE__CLOUDSTORAGEBUCKET"`
	ContentURL         string `json:"contentURL" env:"WEB_TEMPLATE__CONTENTURL"`
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
