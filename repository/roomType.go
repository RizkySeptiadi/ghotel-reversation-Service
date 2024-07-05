package repository

import (
	"hotelrev/database"
	"hotelrev/structs"

	"gorm.io/gorm"
)

type RoomTypeRepository struct {
	db *gorm.DB
}

func NewRoomTypeRepository() *RoomTypeRepository {
	return &RoomTypeRepository{
		db: database.DbConnection,
	}
}

func (repo *RoomTypeRepository) CreateRoomType(roomType *structs.RoomType) error {
	return repo.db.Create(roomType).Error
}

func (repo *RoomTypeRepository) GetRoomTypeByID(id int64) (*structs.RoomType, error) {
	var roomType structs.RoomType
	if err := repo.db.First(&roomType, id).Error; err != nil {
		return nil, err
	}
	return &roomType, nil
}

func (repo *RoomTypeRepository) UpdateRoomType(id int64, updatedRoomType *structs.RoomType) error {
	var roomType structs.RoomType

	// Check if the room type exists
	if err := repo.db.First(&roomType, id).Error; err != nil {
		return err
	}

	// Update the existing room type with the new data
	return repo.db.Model(&roomType).Updates(updatedRoomType).Error
}

func (repo *RoomTypeRepository) DeleteRoomType(id int64) error {
	return repo.db.Delete(&structs.RoomType{}, id).Error
}

func (repo *RoomTypeRepository) GetAllRoomTypes() ([]structs.RoomType, error) {
	var roomTypes []structs.RoomType
	if err := repo.db.Find(&roomTypes).Error; err != nil {
		return nil, err
	}
	return roomTypes, nil
}
