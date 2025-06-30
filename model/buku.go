package model

import "gorm.io/gorm"

type Buku struct {
	ID          int          `gorm:"primaryKey, autoIncrement" json:"id" form:"id"`
	Judul       string       `gorm:"type:text" json:"judul" form:"judul"`
	Penulis     string       `gorm:"type:text" json:"penulis" form:"penulis"`
	TahunTerbit int          `gorm:"type:int" json:"tahun_terbit" form:"tahun_terbit"`
	CategoryID  int          `gorm:"type:int" json:"category_id" form:"category_id"`
	Category    CategoryBuku `gorm:"foreignKey:CategoryID;references:ID" json:"category" form:"category"`
}

func CreateBuku(db *gorm.DB, buku Buku) error {
	result := db.Create(&buku)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func ReadBuku(db *gorm.DB) ([]Buku, error) {
	var buku []Buku
	result := db.Model(&buku).Find(&buku)
	if result.Error != nil {
		return nil, result.Error
	}
	return buku, nil
}

func UpdateBuku(db *gorm.DB, id int, buku Buku) (*Buku, error) {
	var existingBuku Buku
	result := db.First(&existingBuku, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if buku.Judul != "" {
		existingBuku.Judul = buku.Judul
	}
	if buku.Penulis != "" {
		existingBuku.Penulis = buku.Penulis
	}
	if buku.TahunTerbit != 0 {
		existingBuku.TahunTerbit = buku.TahunTerbit
	}

	result = db.Save(&existingBuku)
	if result.Error != nil {
		return nil, result.Error
	}
	return &buku, nil

}

func DeleteBuku(db *gorm.DB, id int) error {
	var buku Buku
	result := db.First(&buku, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&buku)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func GetBukuById(db *gorm.DB, id int) (*Buku, error) {
	var buku Buku
	result := db.First(&buku, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &buku, nil
}
