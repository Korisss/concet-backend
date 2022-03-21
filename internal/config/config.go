package config

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

type Configuration struct {
	Port       int  `json:"port"`
	Production bool `json:"production"`
}

func Load(filePath string) *Configuration {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	var config Configuration

	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	return &config
}
