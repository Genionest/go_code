package one2one

import (
	"fmt"

	"gorm.io/gorm"
)

/* type User struct {
	ID        int
	Name      string
	CompanyID int
	Company   Company // 自动关联 CompanyID 字段(类名+第一个字段名/主键名)
}

type Company struct {
	ID   int
	Name string
}

func Test(db *gorm.DB) {
	db.AutoMigrate(&User{})

	company := Company{
		ID:   1,
		Name: "MicroSoft",
	}
	user := User{
		Name:    "Tom",
		Company: company,
	}
	db.Create(&user)

} */

type User struct {
	ID         int
	Name       string
	CreditCard CreditCard
}

type CreditCard struct {
	ID     int
	Number string
	UserId int
}

func Test(db *gorm.DB) {
	// db.AutoMigrate(&User{})
	db.AutoMigrate(&User{}, &CreditCard{})

	credit_card := CreditCard{
		ID:     1,
		Number: "111",
	}

	user := &User{
		ID:         1,
		Name:       "Tom",
		CreditCard: credit_card,
	}

	db.Create(&user)
}

func Test2(db *gorm.DB) {
	var user User
	db.Preload("CreditCard").First(&user)
	fmt.Println(user)
}
