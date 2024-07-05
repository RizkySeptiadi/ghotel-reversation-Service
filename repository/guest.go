package repository

import (
	"errors"
	"hotelrev/database"
	"hotelrev/structs"
	"time"

	"gorm.io/gorm"
)

type GuestRepository struct {
	db *gorm.DB
}

func NewGuestRepository() *GuestRepository {
	return &GuestRepository{
		db: database.DbConnection,
	}
}
func (repo *GuestRepository) Book(guest *structs.Guest) error {
	// Start a transaction
	tx := repo.db.Begin()

	// Fetch the room with preloaded RoomType by roomID
	var room structs.Room
	if err := tx.Preload("RoomType").First(&room, guest.RoomID).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Check room availability

	// Check room availability
	if room.Availability != 1 {
		tx.Rollback()
		return errors.New("room not available")
	}

	// Populate the guest price with the fetched room price
	guest.Price = room.RoomType.Price
	guest.Total = room.RoomType.Price*float64(guest.Days) - guest.Disc

	// Update room availability to 2
	room.Availability = 2
	if err := tx.Model(&room).Updates(room).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert the guest record into the database
	if err := tx.Create(guest).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}
func (repo *GuestRepository) GetGuestByID(id int64) (*structs.Guest, error) {
	var roomType structs.Guest
	if err := repo.db.First(&roomType, id).Error; err != nil {
		return nil, err
	}
	return &roomType, nil
}

func (repo *GuestRepository) Checkin(id int64, updatedGuest *structs.Guest) error {
	var guest structs.Guest

	// Check if the room type exists
	if err := repo.db.First(&updatedGuest, id).Error; err != nil {
		return err
	}
	if updatedGuest.Status == 1 {

		return errors.New("already check in")
	}
	if updatedGuest.Status == 2 {

		return errors.New("cannot check in checkouted")
	}
	updatedGuest.CheckInAt = time.Now()
	updatedGuest.Status = 1

	// Update the existing room type with the new data
	return repo.db.Model(guest).Where("id = ?", id).Updates(updatedGuest).Error
}
func (repo *GuestRepository) CheckOut(id int64, updatedGuest *structs.Guest) error {
	var guest structs.Guest
	var room structs.Room

	// Find the guest by ID
	if err := repo.db.First(&guest, id).Error; err != nil {
		return err
	}
	if guest.Status == 2 {

		return errors.New("already checkout")
	}

	// Find the room associated with the guest
	if err := repo.db.Model(&room).First(&room, guest.RoomID).Error; err != nil {
		return err
	}

	// Update room availability to 1 (available)
	room.Availability = 1
	if err := repo.db.Save(&room).Error; err != nil {
		return err
	}

	// Update guest data
	updatedGuest.CheckOutAt = time.Now()
	updatedGuest.Status = 2

	// Update the guest with check-out information
	return repo.db.Model(&guest).Updates(updatedGuest).Error
}
func (repo *GuestRepository) Extend(id int64, updatedRoomType *structs.Guest) error {
	var roomType structs.Guest

	// Check if the room type exists
	if err := repo.db.First(&roomType, id).Error; err != nil {
		return err
	}

	// Update the existing room type with the new data
	return repo.db.Model(&roomType).Updates(updatedRoomType).Error
}
func (repo *GuestRepository) DeleteGuest(id int64) error {
	return repo.db.Delete(&structs.Guest{}, id).Error
}

func (repo *GuestRepository) GetAllGuests() ([]structs.Guest, error) {
	var roomTypes []structs.Guest
	if err := repo.db.Find(&roomTypes).Error; err != nil {
		return nil, err
	}
	return roomTypes, nil
}
