package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Config is the config object
type Config struct {
	DB DB
}

// DB is the database config
type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func GetConfig() Config {
	var config Config
	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	return config
}
