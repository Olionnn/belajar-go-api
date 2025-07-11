package model

import "gorm.io/gorm"

type Buku struct {
	ID           int          `gorm:"primaryKey, autoIncrement" json:"id" form:"id"`
	Judul        string       `gorm:"type:text" json:"judul" form:"judul"`
	Penulis      string       `gorm:"type:text" json:"penulis" form:"penulis"`
	TahunTerbit  int          `gorm:"type:int" json:"tahun_terbit" form:"tahun_terbit"`
	CategoryID   int          `gorm:"type:int" json:"category_id" form:"category_id"`
	Category     CategoryBuku `gorm:"foreignKey:CategoryID;references:ID" json:"category" form:"category"`
	RakID        int          `json:"rak_id" form:"rak_id"`
	Rak          Rak          `gorm:"foreignKey:RakID;references:ID" json:"rak" form:"rak"`
	PeminjamanID int          `json:"peminjaman_id" form:"peminjaman_id"`
	Peminjaman   Peminjaman   `gorm:"foreignKey:PeminjamanID;references:ID" json:"peminjaman" form:"peminjaman"`
}

func CreateBuku(db *gorm.DB, buku Buku) error {
	result := db.Create(&buku)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func ReadBuku(db *gorm.DB) ([]Buku, error) {
	var bukuList []Buku
	err := db.Debug().Preload("Rak").Preload("Category").Find(&bukuList).Error
	if err != nil {
		return nil, err
	}
	return bukuList, nil
}

func GetBukuById(db *gorm.DB, id int) (Buku, error) {
	var buku Buku
	err := db.Preload("Rak").First(&buku, id).Error
	if err != nil {
		return Buku{}, err
	}

	response := Buku{
		ID:           buku.ID,
		Judul:        buku.Judul,
		Penulis:      buku.Penulis,
		TahunTerbit:  buku.TahunTerbit,
		RakID:        buku.RakID,
		Rak:          buku.Rak,
		CategoryID:   buku.CategoryID,
		Category:     buku.Category,
		PeminjamanID: buku.PeminjamanID,
		Peminjaman:   buku.Peminjaman,
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

	if updated.Judul != "" {
		buku.Judul = updated.Judul
	}
	if updated.Penulis != "" {
		buku.Penulis = updated.Penulis
	}
	if updated.TahunTerbit != 0 {
		buku.TahunTerbit = updated.TahunTerbit
	}
	if updated.RakID != 0 {
		buku.RakID = updated.RakID
	}
	if updated.CategoryID != 0 {
		buku.CategoryID = updated.CategoryID
	}
	if updated.PeminjamanID != 0 {
		buku.PeminjamanID = updated.PeminjamanID
	}

	err := db.Save(&buku).Error
	return buku, err
}

func DeleteBuku(db *gorm.DB, id int) error {
	return db.Delete(&Buku{}, id).Error
}
