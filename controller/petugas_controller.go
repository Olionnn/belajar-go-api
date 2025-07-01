package controller

import (
	"biodata/database"
	"biodata/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPetugas(c *gin.Context) {

	var petugas []model.Petugas
	var err error

	petugas, err = model.ReadPetugas(database.DB)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": petugas,
	})

}

func CreatePetugas(c *gin.Context) {
	var petugas model.Petugas

	if err := c.ShouldBind(&petugas); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	err := model.CreatePetugas(database.DB, petugas)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Petugas created successfully",
		"data":    petugas,
	})

}

func UpdatePetugas(c *gin.Context) {
	var GetData model.Petugas

	if err := c.ShouldBind(&GetData); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	var id = c.Param("id")
	idConv, err := strconv.Atoi(id)

	petugas := model.Petugas{
		Id:    idConv,
		Nama:  GetData.Nama,
		Jakel: GetData.Jakel,
	}

	hasil, err := model.UpdatePetugas(database.DB, idConv, petugas)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Petugas updated successfully",
		"data":    hasil,
	})

}

func DeletePetugas(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	if id == "" {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	err = model.DeletePetugas(database.DB, idConv)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Petugas deleted successfully",
	})
}

func GetPetugasById(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	petugas, err := model.GetPetugasById(database.DB, idConv)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": petugas,
	})
}
