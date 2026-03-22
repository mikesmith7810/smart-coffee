package main

import (
	"database/sql"
	"log"
	"smart-coffee/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize MySQL connection
	// DSN format: user:password@tcp(service-name:port)/dbname
	// In Kubernetes, MySQL service is accessible via service name DNS
	dsn := "root:coffee-password@tcp(mysql:3306)/coffee_db"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	// Verify connection is working
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to MySQL at mysql:3306")

	// Configure connection pool for observability
	// These settings will be visible in Grafana via metrics
	db.SetMaxOpenConns(25)  // Maximum number of open connections
	db.SetMaxIdleConns(5)   // Maximum number of idle connections
	db.SetConnMaxLifetime(0) // Connections are reused indefinitely

	r := router.New()

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
