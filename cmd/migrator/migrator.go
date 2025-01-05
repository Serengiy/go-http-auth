package main

import (
	"auth_app/internal/config"
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var migrationPath, cfgPath string
	var migrateAction, rollbackAction bool

	flag.StringVar(&migrationPath, "migration-path", "", "Path to a directory containing the migration files")

	flag.StringVar(&cfgPath, "cfg-path", "", "Path to config file")

	flag.BoolVar(&migrateAction, "migrate", false, "Run migration action")
	flag.BoolVar(&rollbackAction, "rollback", false, "Run rollback action")

	flag.Parse()

	cfg := config.MustLoadFromPath(cfgPath)

	if migrationPath == "" {
		panic("migration-path is required")
	}

	if migrateAction == false && rollbackAction == false {
		panic("specify the action: --migrate or --rollback")
	}

	m, err := migrate.New(
		"file://"+migrationPath,
		fmt.Sprintf(
			"%s://%s:%s@localhost:%d/%s?sslmode=%s",
			cfg.Database.Driver, cfg.Database.Username, cfg.Database.Password,
			cfg.Database.Port, cfg.Database.Database, cfg.Database.Sslmode,
		),
	)

	if err != nil {
		panic(err)
	}

	if migrateAction {
		if err := m.Steps(1); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("Nothing to migrate")
			} else {
				fmt.Printf("Error running migrations: %v\n", err)
			}
		} else {
			fmt.Println("Migrations applied")
		}
	}

	if rollbackAction {
		if err := m.Steps(-1); err != nil {
			fmt.Printf("Error running migrations: %v\n", err)
		} else {
			fmt.Println("Database rolled back")
		}
	}
}
