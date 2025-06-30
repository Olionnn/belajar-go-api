package controller

import (
	"biodata/database"
	"biodata/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategoryBuku(c *gin.Context) {
	var categories []model.CategoryBuku
	var err error

	categories, err = model.ReadCategoryBuku(database.DB)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": categories,
	})
}
func CreateCategoryBuku(c *gin.Context) {
	var category model.CategoryBuku

	if err := c.ShouldBind(&category); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	err := model.CreateCategoryBuku(database.DB, category)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Category created successfully",
		"data":    category,
	})
}

func UpdateCategoryBuku(c *gin.Context) {
	var GetData model.CategoryBuku

	if err := c.ShouldBind(&GetData); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID format",
		})
		return
	}

	category, err := model.UpdateCategoryBuku(database.DB, idInt, GetData)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Category updated successfully",
		"data":    category,
	})
}

func DeleteCategoryBuku(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID format",
		})
		return
	}

	err = model.DeleteCategoryBuku(database.DB, idInt)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Category deleted successfully",
	})
}

func GetCategoryBukuByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID format",
		})
		return
	}

	category, err := model.GetCategoryBukuByID(database.DB, idInt)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": category,
	})
}
