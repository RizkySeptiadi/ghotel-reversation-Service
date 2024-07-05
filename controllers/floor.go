package controllers

import (
	"hotelrev/repository"
	"hotelrev/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FloorController struct {
	Repo *repository.FloorRepository
}

func NewFloorController() *FloorController {
	return &FloorController{
		Repo: repository.NewFloorRepository(),
	}
}

func (controller *FloorController) CreateFloor(c *gin.Context) {
	var Floor structs.Floor
	if err := c.ShouldBindJSON(&Floor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.Repo.CreateFloor(&Floor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room type"})
		return
	}

	c.JSON(http.StatusOK, Floor)
}

func (controller *FloorController) GetFloorByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room type ID"})
		return
	}

	Floor, err := controller.Repo.GetFloorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room type not found"})
		return
	}

	c.JSON(http.StatusOK, Floor)
}

func (controller *FloorController) UpdateFloor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room type ID"})
		return
	}

	var updatedFloor structs.Floor
	if err := c.ShouldBindJSON(&updatedFloor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.Repo.UpdateFloor(id, &updatedFloor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room type updated successfully"})
}

func (controller *FloorController) DeleteFloor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room type ID"})
		return
	}

	if err := controller.Repo.DeleteFloor(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete room type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room type deleted successfully"})
}

func (controller *FloorController) GetAllFloors(c *gin.Context) {
	Floors, err := controller.Repo.GetAllFloors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch room types"})
		return
	}

	c.JSON(http.StatusOK, Floors)
}
