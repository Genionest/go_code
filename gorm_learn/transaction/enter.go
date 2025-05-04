package transaction

import (
	"gorm.io/gorm"
)

type TMG struct {
	ID   uint
	Name string
}

func Test(db *gorm.DB) {
	m := db.Migrator()
	if m.HasTable(&TMG{}) {
		m.DropTable(&TMG{})
	}
	m.AutoMigrate(&TMG{})

	// success

	// db.Transaction(func(tx *gorm.DB) error {
	// 	tx.Create(&TMG{Name: "test1"})
	// 	tx.Create(&TMG{Name: "test2"})
	// 	tx.Create(&TMG{Name: "test3"})
	// 	return nil
	// })

	// failed

	// db.Transaction(func(tx *gorm.DB) error {
	// 	tx.Create(&TMG{Name: "test1"})
	// 	tx.Create(&TMG{Name: "test2"})
	// 	tx.Create(&TMG{Name: "test3"})
	// 	return errors.New("error")
	// })

	// nested

	// db.Transaction(func(tx *gorm.DB) error {
	// 	tx.Create(&TMG{Name: "test1"})
	// 	tx.Create(&TMG{Name: "test2"})
	// 	db.Transaction(func(tx *gorm.DB) error {
	// 		tx.Create(&TMG{Name: "test3"})
	// 		return errors.New("error")
	// 	})
	// 	return nil
	// })

	/* manual */

	// success

	// tx := db.Begin()
	// tx.Create(&TMG{Name: "test1"})
	// tx.Create(&TMG{Name: "test2"})
	// tx.Create(&TMG{Name: "test3"})
	// tx.Commit()

	// after commit expression will be ignored

	// tx := db.Begin()
	// tx.Create(&TMG{Name: "test1"})
	// tx.Create(&TMG{Name: "test2"})
	// tx.Commit()
	// tx.Create(&TMG{Name: "test3"})

	// save point and rolback

	tx := db.Begin()
	tx.Create(&TMG{Name: "test1"})
	tx.SavePoint("sp1")
	tx.Create(&TMG{Name: "test2"})
	tx.Create(&TMG{Name: "test3"})
	tx.RollbackTo("sp1") // if transaction has been commited, can't rollback
	tx.Commit()
}
