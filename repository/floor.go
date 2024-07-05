package repository

import (
	"hotelrev/database"
	"hotelrev/structs"

	"gorm.io/gorm"
)

type FloorRepository struct {
	db *gorm.DB
}

func NewFloorRepository() *FloorRepository {
	return &FloorRepository{
		db: database.DbConnection,
	}
}

func (repo *FloorRepository) CreateFloor(roomType *structs.Floor) error {
	return repo.db.Create(roomType).Error
}

func (repo *FloorRepository) GetFloorByID(id int64) (*structs.Floor, error) {
	var roomType structs.Floor
	if err := repo.db.First(&roomType, id).Error; err != nil {
		return nil, err
	}
	return &roomType, nil
}

func (repo *FloorRepository) UpdateFloor(id int64, updatedRoomType *structs.Floor) error {
	var roomType structs.Floor

	// Check if the room type exists
	if err := repo.db.First(&roomType, id).Error; err != nil {
		return err
	}

	// Update the existing room type with the new data
	return repo.db.Model(&roomType).Updates(updatedRoomType).Error
}

func (repo *FloorRepository) DeleteFloor(id int64) error {
	return repo.db.Delete(&structs.Floor{}, id).Error
}

func (repo *FloorRepository) GetAllFloors() ([]structs.Floor, error) {
	var roomTypes []structs.Floor
	if err := repo.db.Find(&roomTypes).Error; err != nil {
		return nil, err
	}
	return roomTypes, nil
}
