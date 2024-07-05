package controllers

import (
	"hotelrev/repository"
	"hotelrev/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GuestController struct {
	Repo     *repository.GuestRepository
	RoomRepo *repository.RoomRepository
}

func NewGuestController() *GuestController {
	return &GuestController{
		Repo: repository.NewGuestRepository(),
	}
}

func (controller *GuestController) Book(c *gin.Context) {
	var Guest structs.Guest
	// var Room structs.Room

	if err := c.ShouldBindJSON(&Guest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Book", "Message": err.Error()})
		return
	}

	// Room, err := controller.RoomRepo.GetRoomByID2(&Guest)
	// if err != nil {

	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
	// 	}
	// }
	// Guest.Price = Room.RoomType.Price
	// Guest.Total = Room.RoomType.Price - Guest.Disc

	if err := controller.Repo.Book(&Guest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Book", "Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Room Successfully Booked"})
}

func (controller *GuestController) CheckOut(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room type ID"})
		return
	}

	var updatedGuest structs.Guest
	if err := c.ShouldBindJSON(&updatedGuest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Checkout", "Message": err.Error()})
		return
	}

	if err := controller.Repo.CheckOut(id, &updatedGuest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Checkout", "Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room type updated successfully"})
}
func (controller *GuestController) CheckIn(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Check in", "Message": err.Error()})
		return
	}

	var updatedGuest structs.Guest
	if err := c.ShouldBindJSON(&updatedGuest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Check in", "Message": err.Error()})
		return
	}

	if err := controller.Repo.Checkin(id, &updatedGuest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Check in", "Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room type updated successfully"})
}

// func (controller *GuestController) get(c *gin.Context) {
// 	Guests, err := controller.Repo.get()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch room types"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, Guests)
// }
