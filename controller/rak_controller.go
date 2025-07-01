package controller

import (
	"biodata/database"
	"biodata/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET all rak
func GetAllRak(c *gin.Context) {
	rakList, err := model.ReadRak(database.DB)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request, invalid data"})
		return
	}
	c.JSON(200, gin.H{"data": rakList})
}

// GET rak by ID
func GetRakByID(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}

	rak, err := model.GetRakById(database.DB, idConv)
	if err != nil {
		c.JSON(400, gin.H{"message": "Rak not found"})
		return
	}
	c.JSON(200, gin.H{"data": rak})
}

// POST create rak
func CreateRak(c *gin.Context) {
	var rak model.Rak

	if err := c.ShouldBind(&rak); err != nil {
		c.JSON(400, gin.H{"message": "Invalid data"})
		return
	}

	err := model.CreateRak(database.DB, rak)
	if err != nil {
		c.JSON(400, gin.H{"message": "Failed to create rak"})
		return
	}

	c.JSON(200, gin.H{"message": "Rak created", "data": rak})
}

// PUT update rak
func UpdateRak(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}

	var newData model.Rak
	if err := c.ShouldBind(&newData); err != nil {
		c.JSON(400, gin.H{"message": "Invalid data"})
		return
	}

	updated, err := model.UpdateRak(database.DB, idConv, newData)
	if err != nil {
		c.JSON(400, gin.H{"message": "Failed to update rak"})
		return
	}

	c.JSON(200, gin.H{"message": "Rak updated", "data": updated})
}

// DELETE rak
func DeleteRak(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}

	err = model.DeleteRak(database.DB, idConv)
	if err != nil {
		c.JSON(400, gin.H{"message": "Failed to delete rak"})
		return
	}

	c.JSON(200, gin.H{"message": "Rak deleted"})
}
