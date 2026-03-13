package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"apirest/internal/infrastructure/persistence/postgres"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Define CLI flags
	upFlag := flag.Bool("up", false, "Apply all up migrations")
	downFlag := flag.Bool("down", false, "Rollback the last migration")
	flag.Parse()

	if !*upFlag && !*downFlag {
		log.Fatal("Please specify an action. Use -up to migrate or -down to rollback.")
	}

	// 2. Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// 3. Construct the Database URL string format required by golang-migrate
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Fallback to building it from your postgres config
		cfg := postgres.LoadConfigFromEnv()
		dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	}

	// 4. Initialize migrator
	log.Println("Connecting to database for migrations...")
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}

	// 5. Execute based on flags
	if *upFlag {
		log.Println("Running UP migrations...")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Could not run up migrations: %v", err)
		}
		log.Println("UP migrations applied successfully!")
	}

	if *downFlag {
		log.Println("Rolling back DOWN one step...")
		// m.Steps(-1) rolls back exactly one migration
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Could not run down migration: %v", err)
		}
		log.Println("DOWN migration applied successfully!")
	}
}
