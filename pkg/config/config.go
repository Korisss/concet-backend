package config

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

type Configuration struct {
	Port       int  `json:"port"`
	Production bool `json:"production"`
	DBConfig   struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		DBName   string `json:"db_name"`
		SSLMode  string `json:"ssl_mode"`
	} `json:"db_config"`
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
