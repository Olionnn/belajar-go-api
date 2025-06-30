package model

type Peminjaman struct {
	ID         int    `gorm:"primaryKey, autoIncrement" json:"id" form:"id"`
	BukuID     int    `json:"buku_id" form:"buku_id"`
	PeminjamID int    `json:"peminjam_id" form:"peminjam_id"`
	Tanggal    string `json:"tanggal" form:"tanggal"`
}

//belum saya buat masi bingungW
