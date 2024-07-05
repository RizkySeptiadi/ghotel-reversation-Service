package database

import (
	"fmt"
	"hotelrev/structs"
	"log"
	"time"

	"gorm.io/gorm"
)

var (
	DbConnection *gorm.DB
)

func DbMigrate(dbParam *gorm.DB) {
	if dbParam == nil {
		log.Fatal("dbParam cannot be nil")
		return
	}
	// DropAllTables(dbParam)

	dbParam.AutoMigrate(&structs.Room{})
	dbParam.AutoMigrate(&structs.Floor{})
	dbParam.AutoMigrate(&structs.Guest{})
	dbParam.AutoMigrate(&structs.RoomType{})
	dbParam.AutoMigrate(&structs.User{})
	dbParam.AutoMigrate(&structs.Role{})

	dbParam.Create(&structs.RoomType{Name: "Regular", Price: 500000})
	dbParam.Create(&structs.RoomType{Name: "Presidential Suite", Price: 1250000})
	role := structs.Role{
		Name:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Create the role record
	if err := dbParam.Create(&role).Error; err != nil {
		fmt.Printf("Failed to create role: %v\n", err)
	} else {
		fmt.Println("Role created successfully!")
	}
	user := structs.User{
		Username:  "admin",
		Password:  "password",
		RoleID:    1, // Assuming you have a role with ID 1 in your database
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Create the user record
	if err := dbParam.Create(&user).Error; err != nil {
		fmt.Printf("Failed to create user: %v\n", err)
	} else {
		fmt.Println("User created successfully!")
	}

}
func Initialize(dbParam *gorm.DB) {
	DbConnection = dbParam
}
func DropAllTables(dbParam *gorm.DB) {
	if dbParam == nil {
		log.Fatal("dbParam cannot be nil")
		return
	}

	dbParam.Migrator().DropTable(&structs.Room{})
	dbParam.Migrator().DropTable(&structs.Floor{})
	dbParam.Migrator().DropTable(&structs.Guest{})
	dbParam.Migrator().DropTable(&structs.RoomType{})
	dbParam.Migrator().DropTable(&structs.User{})
	dbParam.Migrator().DropTable(&structs.Role{})
}
