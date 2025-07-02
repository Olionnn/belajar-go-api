package migration

import (
	"biodata/database"
	"biodata/model"
)

func AutoMigration() error {
	db, err := database.DBConnenction()
	if err != nil {
		return err
	}
	// cara run migration
	// go run main.go migrate
	err = db.AutoMigrate(
		// &model.Petugas{},
		&model.Buku{},
		&model.Rak{},
		&model.CategoryBuku{},
		&model.Peminjaman{},
	)
	if err != nil {
		return err
	}

	return nil
}
