package env

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Server struct with server values
type Server struct {
	Port            uint16 `envconfig:"SERVER_PORT" default:"3000"`
	ShutdownTimeOut uint16 `envconfig:"SERVER_SHUTDOWN_TIMEOUT" default:"10"`
}

// Environment struct with the environment values
type Environment struct {
	Server *Server
}

// LoadEnvironment loads a .env file and set the environment variables
func LoadEnvironment() *Environment {
	err := godotenv.Load()
	if err != nil {
		log.Println("config file not found")
	}

	conf := new(Environment)
	envconfig.Process("", conf)
	return conf
}

// RetrieveEnvVariables retrieve the env variables
func RetrieveEnvVariables() *Environment {
	conf := new(Environment)
	envconfig.Process("", conf)
	return conf
}
