package config

import (
	"inno-lb-master/internal/group"
	"io"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PORT   uint           `yaml:"port"`
	GROUPS []*group.Group `yaml:"groups"`
}

var c *Config
var once sync.Once

func LoadConfig() *Config {

	filename := os.Getenv("CONFIG_PATH")
	if filename == "" {
		filename = "config.yml"
	}

	once.Do(func() {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var buffer []byte
		buffer, err = io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		c = &Config{}
		err = yaml.Unmarshal(buffer, c)
		if err != nil {
			log.Fatal(err)
		}
	})
	return c
}

func ReloadConfig() *Config {
	filename := os.Getenv("CONFIG_PATH")
	if filename == "" {
		filename = "config.yml"
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var buffer []byte
	buffer, err = io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	c = &Config{}
	err = yaml.Unmarshal(buffer, c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
