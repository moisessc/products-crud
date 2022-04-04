package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/golang-migrate/migrate/v4"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"products-crud/pkg/env"
)

// path location of migrations file
const path = "file://database/migrations"

// InitPostgresConnection create an instance of db using the pq library
func InitPostgresConnection(env *env.Database) (*sql.DB, error) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	uri := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(env.User, env.Password),
		Host:     env.Host,
		Path:     env.Name,
		RawQuery: fmt.Sprintf("sslmode=%s", env.SSL),
	}

	retriesTimeOut := time.Duration(env.TimeOut) * time.Second
	timeoutExceeded := time.After(retriesTimeOut)
	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", retriesTimeOut)
		case <-ticker.C:
			db, err := sql.Open("postgres", uri.String())
			if err != nil {
				return nil, fmt.Errorf("could not initialize the db driver: %v", err)
			}

			pingErr := db.Ping()
			if pingErr == nil {
				migrationsRrr := executeMigrations(db, env.Name, 1)
				if migrationsRrr != nil {
					return nil, fmt.Errorf("could not execute migrations: %v", err)
				}
				log.Println(">>> Database online!! <<<")
				return db, nil
			}

			log.Println(fmt.Errorf("could not ping the database: %v", pingErr))
		}
	}
}

// executeMigrations execute the migrations in the database
func executeMigrations(db *sql.DB, databaseName string, version uint) error {
	driver, err := pg.WithInstance(db, &pg.Config{})
	if err != nil {
		return fmt.Errorf("fail to init postgres migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(path, databaseName, driver)
	if err != nil {
		return fmt.Errorf("could not create migration instance: %w", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return fmt.Errorf("could not execute migrations: %w", err)
	}

	log.Println(">>> Migrations executed <<<")

	return nil
}
