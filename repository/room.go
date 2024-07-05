package repository

import (
	"hotelrev/database"
	"hotelrev/structs"

	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{
		db: database.DbConnection,
	}
}

func (repo *RoomRepository) CreateRoom(Room *structs.Room) error {
	return repo.db.Create(Room).Error
}

func (repo *RoomRepository) GetRoomByID(id int64) (*structs.Room, error) {
	var Room structs.Room
	if err := repo.db.First(&Room, id).Error; err != nil {
		return nil, err
	}
	return &Room, nil
}
func (repo *RoomRepository) GetRoomByID2(Guest *structs.Guest) (*structs.Room, error) {
	var Room structs.Room
	if err := repo.db.First(&Room, Guest.RoomID).Error; err != nil {
		return nil, err
	}
	return &Room, nil
}
func (repo *RoomRepository) UpdateRoom(id int64, updatedRoom *structs.Room) error {
	var Room structs.Room

	// Check if the room type exists
	if err := repo.db.First(&Room, id).Error; err != nil {
		return err
	}

	// Update the existing room type with the new data
	return repo.db.Model(&Room).Updates(updatedRoom).Error
}

func (repo *RoomRepository) DeleteRoom(id int64) error {
	return repo.db.Delete(&structs.Room{}, id).Error
}

func (repo *RoomRepository) GetAllRooms() ([]structs.Room, error) {
	var Rooms []structs.Room
	if err := repo.db.Preload("Floor").Preload("RoomType").Preload("Guests").Find(&Rooms).Error; err != nil {
		return nil, err
	}
	return Rooms, nil
}

func (repo *RoomRepository) GetUnBooked() ([]structs.Room, error) {
	var Rooms []structs.Room

	subquery := repo.db.Model(&structs.Guest{}).
		Select("room_id").
		Where("status = ?", 1)

	if err := repo.db.Preload("Floor").
		Preload("RoomType").
		Preload("Guests").
		Not("id IN (?)", subquery).
		Find(&Rooms).Error; err != nil {
		return nil, err
	}

	return Rooms, nil
}
