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
	JwtSigningKey      string `json:"jwtSigningKey" env:"WEB_TEMPLATE__JWTSIGNINGKEY"`
	FireStoreDatabase  string `json:"fireStoreDatabase" env:"WEB_TEMPLATE__FIRESTOREDATABASE"`
	CloudStorageBucket string `json:"cloudStorageBucket" env:"WEB_TEMPLATE__CLOUDSTORAGEBUCKET"`
	ContentURL         string `json:"contentURL" env:"WEB_TEMPLATE__CONTENTURL"`
	GeminiApiKey       string `json:"geminiApiKey" env:"WEB_TEMPLATE__GEMINI_APIKEY"`
}

func LoadConfiguration() {
	var configFile string
	localConfigFile, err := filepath.Abs("config.local.json")
	defaultConfigFile, err := filepath.Abs("config.json")
	_, err = os.Stat(localConfigFile)
	if err == nil {
		configFile = localConfigFile
	} else {
		configFile = defaultConfigFile
	}

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
