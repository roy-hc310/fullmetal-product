package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/roy-hc310/fullmetal-product/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func main() {
	direction := flag.String("direction", "up", "migration direction: up or down")
	steps := flag.Int("steps", 1, "number of steps to migrate (for down/up)")
	flag.Parse()

	err := config.LoadGlobalEnv(".")
	if err != nil {
		fmt.Printf("❌  Failed to load config: %v\n", err)
		return
	}

	err = ensureDBExists()
	if err != nil {
		fmt.Printf("❌  Failed to ensure database exists: %v\n", err)
	}

	err = ensureSchemaExists()
	if err != nil {
		fmt.Printf("❌  Failed to ensure schema exists: %v\n", err)
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s&sslmode=disable",
		config.GlobalEnv.DBWriteUser,
		config.GlobalEnv.DBWritePass,
		config.GlobalEnv.DBWriteHost,
		config.GlobalEnv.DBWritePort,
		config.GlobalEnv.DBWriteName,
		config.GlobalEnv.DBWriteSchema,
	)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("❌  Failed to connect to database: %v\n", err)
	}
	defer db.Close()

	m, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		fmt.Printf("❌  Failed to create migrate instance: %v\n", err)
	}

	switch *direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			fmt.Printf("❌  Migration up failed: %v\n", err)
		}
		fmt.Println("✅ Migration up completed")
	case "down":
		current, dirty, err := m.Version()
		if err != nil && err != migrate.ErrNilVersion {
			fmt.Printf("❌  Failed to get current migration version: %v\n", err)
		}
		if dirty {
			fmt.Println("❌  Database is dirty, cannot migrate down")
			if err := m.Force(int(current)); err != nil {
				fmt.Printf("❌  Failed to force migration version: %v\n", err)
			}
		}
		if err := m.Steps(-*steps); err != nil && err != migrate.ErrNoChange {
			fmt.Printf("❌  Migration down failed: %v\n", err)
		}
		fmt.Printf("✅ Migration down by %d steps completed\n", *steps)
	default:
		fmt.Printf("❌  Invalid migration direction: %s\n", *direction)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Printf("❌  Migration up failed: %v\n", err)
	}

	if srcErr, dbErr := m.Close(); srcErr != nil || dbErr != nil {
		fmt.Printf("❌  Failed to close migration: source error: %v, db error: %v\n", srcErr, dbErr)
	}

	fmt.Println("✅ Migration completed successfully")
}

func ensureDBExists() error {

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		config.GlobalEnv.DBWriteUser,
		config.GlobalEnv.DBWritePass,
		config.GlobalEnv.DBWriteHost,
		config.GlobalEnv.DBWritePort,
	)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("❌  Failed to connect to postgres database: %v\n", err)
	}
	defer db.Close()

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = '%s')", config.GlobalEnv.DBWriteName)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		fmt.Println("Database does not exist. Creating...")
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", config.GlobalEnv.DBWriteName))
		if err != nil {
			return err
		}
	}

	return nil
}

func ensureSchemaExists() error {

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.GlobalEnv.DBWriteUser,
		config.GlobalEnv.DBWritePass,
		config.GlobalEnv.DBWriteHost,
		config.GlobalEnv.DBWritePort,
		config.GlobalEnv.DBWriteName,
	)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("❌  Failed to connect to database: %v\n", err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", config.GlobalEnv.DBWriteSchema))
	if err != nil {
		return err
	}

	return nil
}
