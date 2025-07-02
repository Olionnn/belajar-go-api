package controller

import (
	"biodata/database"
	"biodata/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPeminjaman(c *gin.Context) {
	peminjamanList, err := model.ReadPeminjaman(database.DB)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request"})
		return
	}
	c.JSON(200, peminjamanList)
}

func GetPeminjamanById(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	peminjaman, err := model.GetPeminjamanById(database.DB, idConv)
	if err != nil {
		c.JSON(400, gin.H{"message": "Peminjaman not found"})
		return
	}
	c.JSON(200, peminjaman)
}

func CreatePeminjaman(c *gin.Context) {
	var peminjaman model.Peminjaman
	if err := c.ShouldBind(&peminjaman); err != nil {
		c.JSON(400, gin.H{"message": "Invalid data"})
		return
	}
	err := model.CreatePeminjaman(database.DB, peminjaman)
	if err != nil {
		c.JSON(400, gin.H{"message": "Failed to create peminjaman"})
		return
	}
	c.JSON(200, gin.H{"message": "Peminjaman created", "data": peminjaman})
}

func UpdatePeminjaman(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	var data model.Peminjaman
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"message": "Invalid data"})
		return
	}
	updatedPeminjaman, err := model.UpdatePeminjaman(database.DB, idConv, data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Failed to update peminjaman"})

		return
	}
	c.JSON(200, gin.H{"message": "Peminjaman updated", "data": updatedPeminjaman})
}

func DeletePeminjaman(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	err = model.DeletePeminjaman(database.DB, idConv)
	if err != nil {
		c.JSON(400, gin.H{"message": "Failed to delete peminjaman"})
		return
	}
	c.JSON(200, gin.H{"message": "Peminjaman deleted"})
}
