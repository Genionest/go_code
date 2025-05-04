package one2many

import (
	"gorm.io/gorm"
)

type User struct {
	ID          int
	Name        string
	CreditCards []CreditCard
}

type CreditCard struct {
	ID     int
	Number string
	UserID int
	Bank   Bank
}

type Bank struct {
	ID           int
	CreditCardID int
	Name         string
}

func GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("CreditCards").Find(&users).Error
	return users, err
}

func Test(db *gorm.DB) {
	var user User
	var c1 CreditCard
	var c2 CreditCard

	db.Migrator().DropTable("users", "credit_cards", "banks")
	db.AutoMigrate(&User{}, &CreditCard{}, &Bank{})

	b1 := Bank{
		Name: "CBA",
	}
	b2 := Bank{
		Name: "ICBC",
	}

	c1 = CreditCard{
		Number: "111",
		Bank:   b1,
	}
	c2 = CreditCard{
		Number: "222",
		Bank:   b2,
	}
	user = User{
		Name: "Tom",
		CreditCards: []CreditCard{
			c1,
			c2,
		},
	}

	db.Create(&user)
}

func Test2(db *gorm.DB) {
	// var user User

	// res := db.Preload("CreditCards").First(&user)
	// res = res
	// fmt.Println(user)

	// cards := user.CreditCards
	// fmt.Println(cards)
	// cards[1] = CreditCard{
	// 	Number: "333",
	// }

	// res.Update("CreditCards", cards)

	// fmt.Println(user)

	// db.Preload("CreditCards", "id > ?", "1").First(&user)

	// db.Preload("CreditCards", func(db *gorm.DB) *gorm.DB {
	// 	return db.Where("id > ?", "2")
	// }).First(&user)

	// db.Preload("CreditCards.Bank").Preload("CreditCards").First(&user)

	// db.Preload("CreditCards", func(db *gorm.DB) *gorm.DB {
	// 	return db.Joins("Bank").Where("name = ?", "ICBC")
	// }).First(&user)

	// fmt.Println(user)
}
