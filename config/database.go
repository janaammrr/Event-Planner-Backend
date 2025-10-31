package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the global database handle. It can be nil if connection fails or is skipped.
var DB *gorm.DB

// InitDB tries to initialize a GORM MySQL connection using env variables.
// If any required variable is missing or connection fails, it logs and leaves DB as nil.
func InitDB() {
	user := GetEnv("DB_USER", "")
	pass := GetEnv("DB_PASSWORD", "")
	host := GetEnv("DB_HOST", "")
	port := GetEnv("DB_PORT", "")
	name := GetEnv("DB_NAME", "")

	if user == "" || host == "" || port == "" || name == "" {
		log.Printf("database env not fully set; skipping DB connection")
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return
	}

	DB = db
	log.Printf("database connection initialized")
}
