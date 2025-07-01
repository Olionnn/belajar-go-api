package model

import "gorm.io/gorm"

type Rak struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Nama string `gorm:"type:text" json:"nama" form:"nama"`
}

func ReadRak(db *gorm.DB) ([]Rak, error) {
	var rakList []Rak
	err := db.Find(&rakList).Error
	return rakList, err
}

func GetRakById(db *gorm.DB, id int) (Rak, error) {
	var rak Rak
	err := db.First(&rak, id).Error
	return rak, err
}

func CreateRak(db *gorm.DB, rak Rak) error {
	return db.Create(&rak).Error
}

func UpdateRak(db *gorm.DB, id int, updated Rak) (Rak, error) {
	var rak Rak
	if err := db.First(&rak, id).Error; err != nil {
		return rak, err
	}

	if updated.Nama != "" {
		rak.Nama = updated.Nama
	}

	err := db.Save(&rak).Error
	return rak, err
}

func DeleteRak(db *gorm.DB, id int) error {
	return db.Delete(&Rak{}, id).Error
}
