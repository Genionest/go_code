package association

import "gorm.io/gorm"

type Address struct {
	ID       uint
	Address1 string
	Address2 string
}

type Email struct {
	ID    uint
	Email string
}

type Language struct {
	ID   uint
	Name string
}

type User struct {
	ID                uint
	Name              string
	BillingAddressID  uint
	BillingAddress    Address
	ShippingAddressID uint
	ShippingAddress   Address
	Emails            []Email
	Languages         []Language
}

func Test(db *gorm.DB) {
	m := db.Migrator()
	if m.HasTable(&User{}) {
		m.DropTable(&User{})
	}
	if m.HasTable(&Address{}) {
		m.DropTable(&Address{})
	}
	if m.HasTable(&Email{}) {
		m.DropTable(&Email{})
	}
	if m.HasTable(&Language{}) {
		m.DropTable(&Language{})
	}

	db.AutoMigrate(&User{}, &Address{}, &Email{}, &Language{})

	user := User{
		Name: "Tom",
		BillingAddress: Address{
			Address1: "Billing Address",
		},
		ShippingAddress: Address{
			Address1: "Shipping Address",
		},
		Emails: []Email{
			{
				Email: "111@qq.com",
			},
			{
				Email: "222@qq.com",
			},
		},
		Languages: []Language{
			{
				Name: "English",
			},
			{
				Name: "Chinese",
			},
		},
	}

	db.Create(&user)

	// db.Save(&user)
}
