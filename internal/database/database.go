package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type UriLinkEntry struct {
	Url       string
	Alias     string `gorm:"uniqueIndex"`
	CreatedAt int64
}

func parseEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		msg := fmt.Sprintf("The environment variable '%s' is null.", key)
		panic(msg)
	}

	return value
}

func Init() {

	dbHost := parseEnv("CONCIERGE_PG_HOST")
	dbUser := parseEnv("CONCIERGE_PG_USER")
	dbPass := parseEnv("CONCIERGE_PG_PASS")
	dbName := parseEnv("CONCIERGE_PG_NAME")
	dbPort := parseEnv("CONCIERGE_PG_PORT")

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/London", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to PostgreSQL (via GORM).")
	}

	db.AutoMigrate(&UriLinkEntry{})

	fmt.Println("[Concierge] Connected to PostgreSQL.")

	DB = db
}
