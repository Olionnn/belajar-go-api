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

	err = db.AutoMigrate(
		// &model.Petugas{},
		// &model.Buku{},
		&model.Rak{},
	)
	if err != nil {
		return err
	}

	return nil
}
