package controllers

import (
	"hotelrev/repository"
	"hotelrev/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	Repo *repository.RoomRepository
}

func NewRoomController() *RoomController {
	return &RoomController{
		Repo: repository.NewRoomRepository(),
	}
}

func (controller *RoomController) Create(c *gin.Context) {
	var Room structs.Room
	if err := c.ShouldBindJSON(&Room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.Repo.CreateRoom(&Room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room "})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "New Room Created"})
}

func (controller *RoomController) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room  ID"})
		return
	}

	Room, err := controller.Repo.GetRoomByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room  not found"})
		return
	}

	c.JSON(http.StatusOK, Room)
}

func (controller *RoomController) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room  ID"})
		return
	}

	var updatedRoom structs.Room
	if err := c.ShouldBindJSON(&updatedRoom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.Repo.UpdateRoom(id, &updatedRoom); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room "})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room  updated successfully"})
}

func (controller *RoomController) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room  ID"})
		return
	}

	if err := controller.Repo.DeleteRoom(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete room "})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room  deleted successfully"})
}

func (controller *RoomController) Get(c *gin.Context) {
	Rooms, err := controller.Repo.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Rooms "})
		return
	}

	c.JSON(http.StatusOK, Rooms)
}

func (controller *RoomController) GetUnBooked(c *gin.Context) {
	Rooms, err := controller.Repo.GetUnBooked()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Rooms "})
		return
	}

	c.JSON(http.StatusOK, Rooms)
}
