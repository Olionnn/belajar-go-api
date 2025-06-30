package model

import "gorm.io/gorm"

type CategoryBuku struct {
	ID   int    `gorm:"primaryKey, autoIncrement" json:"id" form:"id"`
	Nama string `gorm:"type:text" json:"nama" form:"nama"`
}

func CreateCategoryBuku(db *gorm.DB, category CategoryBuku) error {
	if result := db.Create(&category); result.Error != nil {
		return result.Error
	}
	return nil
}

func ReadCategoryBuku(db *gorm.DB) ([]CategoryBuku, error) {
	var categories []CategoryBuku
	if result := db.Model(&categories).Find(&categories); result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func UpdateCategoryBuku(db *gorm.DB, id int, category CategoryBuku) (*CategoryBuku, error) {
	var existingCategory CategoryBuku
	if result := db.First(&existingCategory, id); result.Error != nil {
		return nil, result.Error
	}

	if category.Nama != "" {
		existingCategory.Nama = category.Nama
	}

	if result := db.Save(&existingCategory); result.Error != nil {
		return nil, result.Error
	}
	return &existingCategory, nil

}

func DeleteCategoryBuku(db *gorm.DB, id int) error {
	var category CategoryBuku
	if result := db.First(&category, id); result.Error != nil {
		return result.Error
	}

	if result := db.Delete(&category); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCategoryBukuByID(db *gorm.DB, id int) (*CategoryBuku, error) {
	var category CategoryBuku
	if result := db.First(&category, id); result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}
