package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"api-bootstrap-echo/libs/logger"
	"api-bootstrap-echo/models"

	"github.com/kelseyhightower/envconfig"
)

var configuration models.Configuration

// NewConfiguration :
func NewConfiguration() *models.Configuration {
	configuration = models.Configuration{}
	readConfigFile()
	readConfigEnv()
	logger.Debug("config", "NewConfiguration", fmt.Sprintf("%+v", configuration))
	return &configuration
}

func readConfigFile() {
	var _, b, _, _ = runtime.Caller(0)
	var basepath = filepath.Dir(b)
	var defaultConfigFile = basepath + "/config.default.json"
	file, err := os.Open(defaultConfigFile)
	if err != nil {
		logger.Fatal("config", "readConfigFile", defaultConfigFile, err.Error())
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		logger.Fatal("config", "readConfigFile", "decode", err.Error())
	}
}

func readConfigEnv() {
	err := envconfig.Process("", &configuration)
	if err != nil {
		logger.Fatal("config", "readConfigEnv", "environmeent", err.Error())
	}
}
