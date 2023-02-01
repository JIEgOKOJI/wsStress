package config

import (
	"chatLoadTest/data"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Path string
}

func (self *Config) Load() data.ParsedConfig {
	log.Println("Loading config ", self.Path)
	f, err := os.ReadFile(self.Path)
	if err != nil {
		log.Fatal("Error openning config file", err)
	}
	parsed := self.parse(f)
	return parsed
}

func (self *Config) parse(configFile []byte) data.ParsedConfig {
	var c data.ParsedConfig
	if err := yaml.Unmarshal(configFile, &c); err != nil {
		log.Fatal("error parsing config ", err)
	}
	return c
}
