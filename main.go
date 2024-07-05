package main

import (
	"database/sql"
	"fmt"
	"hotelrev/database"
	"hotelrev/middleware"
	"hotelrev/routers"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *sql.DB
	err error
)

// const (
// 	dbHost     = "localhost"
// 	dbPort     = 5432
// 	dbUser     = "postgres"
// 	dbPassword = "2209"
// 	dbName     = "hotel"
// )

func main() {

	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("ENV Failed to load!")
		panic(err)
	} else {
		fmt.Println("ENV Success !")

	}
	dbHost := os.Getenv("PGHOST")
	dbPort, _ := strconv.Atoi(os.Getenv("PGPORT"))
	dbUser := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")

	// Create connection string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB Connection Failed !")
		panic(err)
	} else {
		fmt.Println("DB Connection Success !")

	}
	// fmt.Println(db)
	database.DbMigrate(db)
	database.Initialize(db)

	middleware.SetDatabase(db)

	// defer DB.Close()
	PORT := ":8000"
	routers.StartServer().Run(PORT)

}
