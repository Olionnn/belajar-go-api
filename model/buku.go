package model

import "gorm.io/gorm"

type Buku struct {
	ID          int          `gorm:"primaryKey, autoIncrement" json:"id" form:"id"`
	Judul       string       `gorm:"type:text" json:"judul" form:"judul"`
	Penulis     string       `gorm:"type:text" json:"penulis" form:"penulis"`
	TahunTerbit int          `gorm:"type:int" json:"tahun_terbit" form:"tahun_terbit"`
	CategoryID  int          `gorm:"type:int" json:"category_id" form:"category_id"`
	Category    CategoryBuku `gorm:"foreignKey:CategoryID;references:ID" json:"category" form:"category"`
	RakID       int          `gorm:"not null"`
	Rak         Rak          `gorm:"foreignKey=RakID"`
}

func CreateBuku(db *gorm.DB, buku Buku) error {
	result := db.Create(&buku)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func ReadBuku(db *gorm.DB) ([]BukuResponse, error) {
	var bukuList []Buku
	err := db.Preload("Rak").Find(&bukuList).Error
	if err != nil {
		return nil, err
	}

	var response []BukuResponse
	for _, b := range bukuList {
		response = append(response, BukuResponse{
			ID:          b.ID,
			Judul:       b.Judul,
			Penulis:     b.Penulis,
			TahunTerbit: b.TahunTerbit,
			RakID:       b.RakID,
			RakNama:     b.Rak.Nama,
		})
	}

	return response, nil
}

func GetBukuById(db *gorm.DB, id int) (BukuResponse, error) {
	var buku Buku
	err := db.Preload("Rak").First(&buku, id).Error
	if err != nil {
		return BukuResponse{}, err
	}

	response := BukuResponse{
		ID:          buku.ID,
		Judul:       buku.Judul,
		Penulis:     buku.Penulis,
		TahunTerbit: buku.TahunTerbit,
		RakID:       buku.RakID,
		RakNama:     buku.Rak.Nama,
	}

	return response, nil
}

// func CreateBuku(db *gorm.DB, buku Buku) error {
// 	return db.Create(&buku).Error
// }

func UpdateBuku(db *gorm.DB, id int, updated Buku) (Buku, error) {
	var buku Buku
	if err := db.First(&buku, id).Error; err != nil {
		return buku, err
	}
	buku.Judul = updated.Judul
	buku.Penulis = updated.Penulis
	buku.TahunTerbit = updated.TahunTerbit
	buku.RakID = updated.RakID
	err := db.Save(&buku).Error
	return buku, err
}

func DeleteBuku(db *gorm.DB, id int) error {
	return db.Delete(&Buku{}, id).Error
}
