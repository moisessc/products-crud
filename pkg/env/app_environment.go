package env

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	blankPrefix = ""
)

// Server struct with server values
type Server struct {
	Port            uint16 `envconfig:"SERVER_PORT" default:"3000"`
	ShutdownTimeOut uint16 `envconfig:"SERVER_SHUTDOWN_TIMEOUT" default:"10"`
}

// Database struct with database environment values
type Database struct {
	Host     string `envconfig:"DATABASE_HOST" default:"localhost"`
	Name     string `envconfig:"DATABASE_NAME"`
	User     string `envconfig:"DATABASE_USER"`
	Password string `envconfig:"DATABASE_PASSWORD"`
	SSL      string `envconfig:"DATABASE_SSL_MODE" default:"disable"`
	TimeOut  int64  `envconfig:"DATABASE_TIME_OUT" default:"60"`
}

// Environment struct with the environment values
type Environment struct {
	Server   *Server
	Database *Database
}

// LoadEnvironment loads a .env file and set the environment variables
func LoadEnvironment() *Environment {
	err := godotenv.Load()
	if err != nil {
		log.Println("config file not found")
	}

	conf := new(Environment)
	envconfig.Process(blankPrefix, conf)
	return conf
}

// RetrieveEnvVariables retrieve the env variables
func RetrieveEnvVariables() *Environment {
	conf := new(Environment)
	envconfig.Process(blankPrefix, conf)
	return conf
}
