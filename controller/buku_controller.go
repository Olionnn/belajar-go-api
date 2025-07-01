package controller

import (
	"biodata/database"
	"biodata/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBuku(c *gin.Context) {
	bukuList, err := model.ReadBuku(database.DB)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request"})
		return
	}
	c.JSON(200, gin.H{"data": bukuList})
}

func GetBukuById(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	buku, err := model.GetBukuById(database.DB, idConv)
	if err != nil {
		c.JSON(400, gin.H{"message": "Buku not found"})
		return
	}
	c.JSON(200, gin.H{"data": buku})
}

func CreateBuku(c *gin.Context) {
	var buku model.Buku
	if err := c.ShouldBind(&buku); err != nil {
		c.JSON(400, gin.H{"message": "Invalid data"})
		return
	}
	err := model.CreateBuku(database.DB, buku)
	if err != nil {
		c.JSON(400, gin.H{"message": "Gagal membuat buku"})
		return
	}
	c.JSON(200, gin.H{"message": "Buku created", "data": buku})
}

func UpdateBuku(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	var data model.Buku
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"message": "Invalid data"})
		return
	}
	updated, err := model.UpdateBuku(database.DB, idConv, data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Failed to update buku"})
		return
	}
	c.JSON(200, gin.H{"message": "Updated", "data": updated})
}

func DeleteBuku(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	err = model.DeleteBuku(database.DB, idConv)
	if err != nil {
		c.JSON(400, gin.H{"message": "Failed to delete buku"})
		return
	}
	c.JSON(200, gin.H{"message": "Deleted"})
}
