package model

import "gorm.io/gorm"

type Petugas struct {
	Id    int    `gorm:"primaryKey, autoIncrement" json:"id" form:"id"`
	Nama  string `gorm:"type:text, not null" json:"nama" form:"nama"`
	Jakel string `gorm:"type:text, not null" json:"jakel" form:"jakel"`
}

func CreatePetugas(db *gorm.DB, petugas Petugas) error {
	result := db.Create(&petugas)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func ReadPetugas(db *gorm.DB) ([]Petugas, error) {
	var petugas []Petugas
	result := db.Model(&petugas).Find(&petugas)
	if result.Error != nil {
		return nil, result.Error
	}
	return petugas, nil
}

func UpdatePetugas(db *gorm.DB, id int, petugas Petugas) (*Petugas, error) {
	var existingPetugas Petugas
	result := db.First(&existingPetugas, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if petugas.Nama != "" {
		existingPetugas.Nama = petugas.Nama
	}
	if petugas.Jakel != "" {
		existingPetugas.Jakel = petugas.Jakel
	}

	result = db.Save(&existingPetugas)
	if result.Error != nil {
		return nil, result.Error
	}
	return &petugas, nil

}

func DeletePetugas(db *gorm.DB, id int) error {
	var petugas Petugas
	result := db.First(&petugas, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&petugas)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func GetPetugasById(db *gorm.DB, id int) (*Petugas, error) {
	var petugas Petugas
	result := db.First(&petugas, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &petugas, nil
}
