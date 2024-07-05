package controllers

import (
	"hotelrev/repository"
	"hotelrev/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomTypeController struct {
	Repo *repository.RoomTypeRepository
}

func NewRoomTypeController() *RoomTypeController {
	return &RoomTypeController{
		Repo: repository.NewRoomTypeRepository(),
	}
}

func (controller *RoomTypeController) CreateRoomType(c *gin.Context) {
	var roomType structs.RoomType
	if err := c.ShouldBindJSON(&roomType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.Repo.CreateRoomType(&roomType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Room  Type Successfully created"})
}

func (controller *RoomTypeController) GetRoomTypeByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room type ID"})
		return
	}

	roomType, err := controller.Repo.GetRoomTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room type not found"})
		return
	}

	c.JSON(http.StatusOK, roomType)
}

func (controller *RoomTypeController) UpdateRoomType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room type ID"})
		return
	}

	var updatedRoomType structs.RoomType
	if err := c.ShouldBindJSON(&updatedRoomType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.Repo.UpdateRoomType(id, &updatedRoomType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room type updated successfully"})
}

func (controller *RoomTypeController) DeleteRoomType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room type ID"})
		return
	}

	if err := controller.Repo.DeleteRoomType(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete room type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room type deleted successfully"})
}

func (controller *RoomTypeController) GetAllRoomTypes(c *gin.Context) {
	roomTypes, err := controller.Repo.GetAllRoomTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch room types"})
		return
	}

	c.JSON(http.StatusOK, roomTypes)
}
