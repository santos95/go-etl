package config

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
)

var fileName string = "./config/config.yaml"

type config struct {
	Server string `yaml:"server"`
	Port string `yaml:"port"`
	Database string `yaml: database`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	MongoURI string `yaml:"mongouri"`
	DatabaseMongo string `yaml: "databasemongo"`
	Collection string `yaml: "collection"`
	Batchsize string `yaml:"batchsize"`
}

// Read and parse the config.yaml file
func GetConfigValues() (*config, error){


	// Read the YAML file 
	dat, err := os.ReadFile(fileName)
	if err != nil {

		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	// variable to hold the structure parameters
	var c config 

	// unmarshal the content of dat into c
	err = yaml.Unmarshal(dat, &c)
	if err != nil {

		return nil, fmt.Errorf("Error unmarshalling YAML: %v", err)
	}

	return &c, nil 
}