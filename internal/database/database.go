package database

import (
	"fmt"
	"os"

	"github.com/abiddiscombe/concierge/internal/log"
	slogGorm "github.com/alfonmga/slog-gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type UriLinkEntry struct {
	gorm.Model
	Url                 string
	Alias               string `gorm:"uniqueIndex"`
	CreatedAt           int64
	UriActivationEvents []UriActivationEvent
}

type UriActivationEvent struct {
	gorm.Model
	UriLinkEntryId uint
	IsRedirect     bool
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

	logger := log.NewLogger("database")

	DBLogger := slogGorm.New(
		slogGorm.WithLogger(logger),
	)

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=prefer TimeZone=Europe/London", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: DBLogger,
	})

	if err != nil {
		msg := "Failed to connect to PostgreSQL"
		logger.Error(msg)
		panic(msg)
	}

	err = db.AutoMigrate(&UriLinkEntry{}, &UriActivationEvent{})

	if err != nil {
		msg := "Failed to sync models with PostgreSQL"
		logger.Error(msg)
		panic(msg)
	}

	logger.Info("Connected to PostgreSQL")
	DB = db
}
