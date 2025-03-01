package config

import (
	"io"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PORT       uint   `yaml:"port"`
	ALIAS      string `yaml:"alias"`
	MASTER_URL string `yaml:"master_url"`
}

var c *Config
var once sync.Once

func LoadConfig() *Config {
	const op = "pkg.config.load"

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
