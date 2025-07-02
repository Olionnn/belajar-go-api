package model

import "gorm.io/gorm"

type Peminjaman struct {
	ID     int  `gorm:"primaryKey, autoIncrement" json:"id" form:"id"`
	BukuID int  `json:"buku_id" form:"buku_id"`
	Buku   Buku `gorm:"foreignKey:BukuID;references:ID" json:"buku" form:"buku"`
	// PeminjamID is the ID of the person borrowing the book
	//jadi saya
	PeminjamID   int    `json:"peminjam_id" form:"peminjam_id"`
	NamaPeminjam string `gorm:"type:text, not null" json:"nama_peminjam" form:"nama_peminjam"`
	Tanggal      string `json:"tanggal" form:"tanggal"`
}

func CreatePeminjaman(db *gorm.DB, peminjaman Peminjaman) error {
	result := db.Create(&peminjaman)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ReadPeminjaman(db *gorm.DB) ([]Peminjaman, error) {
	var peminjamanList []Peminjaman
	err := db.Find(&peminjamanList).Error
	if err != nil {
		return nil, err
	}

	return peminjamanList, nil
}

func GetPeminjamanById(db *gorm.DB, id int) (Peminjaman, error) {
	var peminjaman Peminjaman
	err := db.First(&peminjaman, id).Error
	if err != nil {
		return Peminjaman{}, err
	}

	return peminjaman, nil
}

func UpdatePeminjaman(db *gorm.DB, id int, peminjaman Peminjaman) (*Peminjaman, error) {
	var existingPeminjaman Peminjaman
	if result := db.First(&existingPeminjaman, id); result.Error != nil {
		return nil, result.Error
	}

	if peminjaman.BukuID != 0 {
		existingPeminjaman.BukuID = peminjaman.BukuID
	}
	if peminjaman.PeminjamID != 0 {
		existingPeminjaman.PeminjamID = peminjaman.PeminjamID
	}
	if peminjaman.Tanggal != "" {
		existingPeminjaman.Tanggal = peminjaman.Tanggal
	}

	if result := db.Save(&existingPeminjaman); result.Error != nil {
		return nil, result.Error
	}
	return &existingPeminjaman, nil
}

func DeletePeminjaman(db *gorm.DB, id int) error {
	var peminjaman Peminjaman
	if result := db.First(&peminjaman, id); result.Error != nil {
		return result.Error
	}

	if result := db.Delete(&peminjaman); result.Error != nil {
		return result.Error
	}
	return nil
}
func GetPeminjamanByBukuID(db *gorm.DB, bukuID int) ([]Peminjaman, error) {
	var peminjamanList []Peminjaman
	err := db.Where("buku_id = ?", bukuID).Find(&peminjamanList).Error
	if err != nil {
		return nil, err
	}

	return peminjamanList, nil
}
