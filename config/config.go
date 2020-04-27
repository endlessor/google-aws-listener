package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var Cfg *Config = &Config{}

type Config struct {
	ServerConfigurations ServerConfigurations
	FileUploader         FileUploader
	AWS                  AWS
	Logger               Logger
}

type ServerConfigurations struct {
	Port         string
	InstanceName string
	WorkingDir   string
}

type FileUploader struct {
	ChunkSize       int64
	LocalStorageDir string
}

type AWS struct {
	Region string
	S3     S3
}

type S3 struct {
	Bucket string
}

type Logger struct {
	FileLocation string
	Level        string
}

func LoadConfig() {
	workingDir := os.Getenv("WORKING_DIR")

	configFile := fmt.Sprintf("%s%s", workingDir, "config/config.json")

	file, errOpenFile := os.Open(configFile)

	if errOpenFile != nil {
		log.Fatal(errOpenFile)
	}

	decoder := json.NewDecoder(file)

	configuration := Config{}
	err := decoder.Decode(&configuration)

	configuration.ServerConfigurations.WorkingDir = workingDir
	configuration.FileUploader.LocalStorageDir = fmt.Sprintf("%s%s", workingDir, configuration.FileUploader.LocalStorageDir)
	configuration.Logger.FileLocation = fmt.Sprintf("%s%s", workingDir, configuration.Logger.FileLocation)

	if err != nil {
		log.Fatal(err)
	}

	Cfg = &configuration
}
