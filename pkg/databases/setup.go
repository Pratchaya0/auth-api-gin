// package databases

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/Pratchaya0/auth-api-gin/configs"
// 	"github.com/Pratchaya0/auth-api-gin/pkg/databases/providers"
// 	"github.com/joho/godotenv"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func DB() *gorm.DB {
// 	return db
// }

// func SetupDatabase() {

// 	if err := godotenv.Load("../.env"); err != nil {
// 		panic(err.Error())
// 	}

// 	cfg := new(configs.Configs)

// 	// Fiber configs
// 	cfg.App.Host = os.Getenv("GIN_HOST")
// 	cfg.App.Port = os.Getenv("GIN_PORT")

// 	// Database Configs
// 	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
// 	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
// 	cfg.PostgreSQL.Protocol = os.Getenv("DB_PROTOCOL")
// 	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
// 	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
// 	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")

// 	database, err := providers.NewPostgresDBConnection(cfg)
// 	if err != nil {
// 		log.Printf("error, can't connect to database, %s", err.Error())
// 	}

// 	database.AutoMigrate()

// 	db = database

//		fmt.Println("Database migration completed!")
//	}
package databases

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
