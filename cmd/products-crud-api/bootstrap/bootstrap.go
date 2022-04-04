package bootstrap

import (
	"fmt"

	"products-crud/pkg/env"
)

// Run retrieves the environment, builds the server router and starts the server
func Run() error {
	appEnv := env.LoadEnvironment()

	router := newEchoRouter()
	srv := newServer(appEnv.Server, router)

	err := srv.up()
	if err != nil {
		return fmt.Errorf("failed to init server, %v", err)
	}
	return nil
}
