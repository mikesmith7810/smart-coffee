package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"smart-coffee/config"
	"smart-coffee/handlers"
	"smart-coffee/repository"
	"smart-coffee/router"
	"smart-coffee/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configPath := "config.yaml"
	if value := os.Getenv("SMART_COFFEE_CONFIG"); value != "" {
		configPath = value
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Printf("Successfully connected to MySQL at %s:%d", cfg.Database.Host, cfg.Database.Port)

	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	db.SetConnMaxLifetime(0)

	if err := initSchema(db); err != nil {
		log.Fatalf("Failed to initialise schema: %v", err)
	}

	repo := repository.NewCoffeeRepository(db)
	svc := service.NewCoffeeService(repo)
	h := handlers.NewHandler(svc)
	r := router.New(h)

	log.Printf("Server starting on :%s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func initSchema(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS coffees (
			id       VARCHAR(255) PRIMARY KEY,
			name     VARCHAR(255) NOT NULL,
			calories INT         NOT NULL
		)
	`)
	return err
}
