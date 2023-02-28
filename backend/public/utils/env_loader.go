package utils

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Env map[string]string
}

func (c *Config) LoadEnv() {
	loadedEnv, err := godotenv.Read()

	if err != nil {
		log.Fatalln("Failed to load env file.")
	}

	c.Env = loadedEnv

}
