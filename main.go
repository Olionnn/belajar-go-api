package main

import (
	"biodata/database"
	"biodata/migration"
	"biodata/route"
	"os"
)

func main() {
	//command migrate

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		err := migration.AutoMigration()
		if err != nil {
			panic("Gagal melakukan migrasi: " + err.Error())
		}
		println("Migrasi berhasil")
		return
	}

	_, err := database.DBConnenction()
	if err != nil {
		panic("Gagal koneksi ke database: " + err.Error())
	}

	route := route.Route()

	route.Run(":2222")

	// route.Run(":8888")

}

// import "fmt"

// type Buku struct {
// 	Judul       string
// 	Penulis     string
// 	TahunTerbit int
// }

// type Petugas struct {
// 	Id    int
// 	Nama  string
// 	Jakel string
// }

// func main() {

// 	variable1 := 9
// 	variable2 := 9

// 	arrString := []string{}

// 	result := tambah(variable1, variable2)
// 	fmt.Println("Hasil penjumlahan:", result)

// 	arrString = isiArray(arrString, "neva")
// 	arrString = isiArray(arrString, "Haydar")

// 	fmt.Println("Isi array:", arrString)

// 	buku := Buku{} // 	0x0

// 	fmt.Println("Buku sebelum diubah:", buku)

// 	buku = Buku{
// 		Judul:       "Belajar Go",
// 		Penulis:     "John Doe",
// 		TahunTerbit: 2023,
// 	}
// 	addBuku(buku, "Belajar Go Lanjutan", "judul")
// 	fmt.Println("Buku setelah diubah:", buku)

// hasil := cothSwitch(4)
// fmt.Println("Hasil switch:", hasil)

// func tambah(param1 int, param2 int) int {
// 	result := param1 + param2
// 	return result
// }

// func isiArray(arr []string, value string) []string {
// 	arr = append(arr, value)
// 	return arr
// }

// func addBuku(buku Buku, value string, field string) {
// 	switch field {
// 	case "judul":
// 		buku.Judul = value
// 	case "penulis":
// 		buku.Penulis = value
// 	case "tahunTerbit":
// 		fmt.Sscanf(value, "%d", &buku.TahunTerbit)
// 	default:
// 		fmt.Println("Field tidak dikenal")
// 	}
// }

// func cothSwitch(value int) string {
// 	switch value {
// 	case 1:
// 		return "Satu"
// 	case 2:
// 		return "Dua"
// 	case 3:
// 		return "Tiga"
// 	default:
// 		return "Tidak diketahui"
// 	}

// 	// if value == 1 {
// 	// 	return "Satu"
// 	// } else if value == 2 {
// 	// 	return "Dua"
// 	// } else if value == 3 {
// 	// 	return "Tiga"
// 	// } else {
// 	// 	return "Tidak diketahui"
// 	// }
// }

// func CreatePetugas(petugas Petugas, id int, nama string, jakel string) Petugas {
// 	petugas.Id = id
// 	petugas.Nama = nama
// 	petugas.Jakel = jakel
// 	return petugas
// }

// func ReadPetugas(petugas Petugas) {
// 	fmt.Println("ID:", petugas.Id)
// 	fmt.Println("Nama:", petugas.Nama)
// 	fmt.Println("Jenis Kelamin:", petugas.Jakel)
// }

// func UpdatePetugas(petugas *Petugas, id int, nama string, jakel string) {
// 	petugas.Id = id
// 	petugas.Nama = nama
// 	petugas.Jakel = jakel
// }

// func DeletePetugas(petugas *Petugas) {
// 	petugas.Id = 0
// 	petugas.Nama = ""
// 	petugas.Jakel = ""
// 	fmt.Println("Petugas telah dihapus")
// }

// func SearchPetugas(petugas []Petugas, id int) *Petugas {
// 	for i := range petugas {
// 		if petugas[i].Id == id {
// 			return &petugas[i]
// 		}
// 	}
// 	return nil
// }
