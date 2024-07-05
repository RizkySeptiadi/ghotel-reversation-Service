package structs

import (
	"time"

	"gorm.io/gorm"
)

// Room struct with relationships
type Room struct {
	gorm.Model

	ID           int64
	Number       int64  `gorm:"uniqueIndex:idx_number"`
	Section      string `gorm:"uniqueIndex:idx_number"`
	FloorID      int64  `gorm:"uniqueIndex:idx_number"` // Foreign key for Floor
	TypeID       int64  // Foreign key for RoomType
	Availability int64
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Floor    Floor    `gorm:"foreignKey:FloorID"` // Belongs to Floor
	RoomType RoomType `gorm:"foreignKey:TypeID"`  // Belongs to RoomType
	Guests   []Guest  `gorm:"foreignKey:RoomID"`  // Has many Guests
}

// RoomType struct with relationships
type RoomType struct {
	gorm.Model

	ID        int64
	Name      string `gorm:"unique"`
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time

	Rooms []Room `gorm:"foreignKey:TypeID"` // Has many Rooms
}

// Floor struct with relationships
type Floor struct {
	gorm.Model

	ID        int64
	Floor     int64 `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Rooms []Room `gorm:"foreignKey:FloorID"` // Has many Rooms
}

// Guest struct with relationships
type Guest struct {
	gorm.Model

	ID         int64
	RoomID     int64 // Foreign key for Room
	FirstName  string
	LastName   string
	Phone      string
	Address    string
	Price      float64
	Disc       float64
	Total      float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CheckInAt  time.Time
	CheckOutAt time.Time
	Status     int64
	Days       int64

	Room Room `gorm:"foreignKey:RoomID"`
}
type User struct {
	gorm.Model

	ID        int64
	Username  string `gorm:"unique"`
	Password  string
	RoleID    int64
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role struct {
	gorm.Model

	ID        int64
	Name      string `gorm:"unique"`
	Users     []User `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
