package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"    // or the Docker service name if running in another container
	port     = 5432           // default PostgreSQL port
	user     = "myuser"       // as defined in docker-compose.yml
	password = "mypassword"   // as defined in docker-compose.yml
	dbname   = "auth-api-gin" // as defined in docker-compose.yml
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	//Migrate the schema
	database.AutoMigrate(
	// Add schema
	)

	// Assign to global variable
	db = database

	fmt.Println("Database postgres migration completed!")
}
