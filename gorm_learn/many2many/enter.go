package many2many

import (
	"gorm.io/gorm"
)

type User struct {
	ID uint
	// gorm.Model
	Languages []Language `gorm:"many2many:user_languages"`
}

type Language struct {
	ID uint
	// gorm.Model
	Users []User `gorm:"many2many:user_languages"`
}

func Test(db *gorm.DB) {
	m := db.Migrator()
	if m.HasTable(&User{}) {
		m.DropTable(&User{})
	}
	if m.HasTable(&Language{}) {
		m.DropTable(&Language{})
	}
	if m.HasTable("user_languages") {
		m.DropTable("user_languages")
	}

	m.AutoMigrate(&User{})

	L1 := Language{}
	L2 := Language{}
	L3 := Language{}
	LL := []Language{L1, L2}

	u1 := User{Languages: LL}
	u2 := User{Languages: LL}
	u3 := User{}

	// follow incorret
	// u1 := User{Languages: []Language{L1, L2}}
	// u2 := User{Languages: []Language{L1, L2}}
	// that will create four Language (ID: 1-4)

	db.Create(&u1)
	db.Create(&u2)
	db.Create(&u3)
	// auto-mate fill id, if the line in first, L3.ID will be 1
	db.Create(&L3)
}

func Test2(db *gorm.DB) {
	var user User
	var lan Language
	user = user
	lan = lan

	// // result is equal
	// // be care of that use Language(s), it's field name of User
	// db.Model(&User{}).Preload("Languages").Find(&user, 1)
	// db.Preload("Languages").Find(&user, 1)

	// var ls []Language
	// user = User{ID: 1}
	// db.Model(&user).Preload("Users").Association("Languages").Find(&lan, 2)
	/* SELECT `languages`.`id` FROM `languages` JOIN `user_languages`
	ON `user_languages`.`language_id` = `languages`.`id`
	AND `user_languages`.`user_id` = 1 WHERE `languages`.`id` = 2 */

	user = User{ID: 3}
	// db.Model(&user).Association("Languages").Append(&Language{ID: 3}, &Language{ID: 2})
	// db.Model(&user).Association("Languages").Delete(&Language{ID: 3})
	// // Replace is delete all and append new
	db.Model(&user).Association("Languages").Replace(&Language{ID: 2}, &Language{ID: 3})
	// db.Model(&user).Association("Languages").Clear()

	// fmt.Println(lan)
}
