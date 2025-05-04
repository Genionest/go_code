package polymorphic

import "gorm.io/gorm"

type Pig struct {
	ID   uint
	Name string
	// if don't auto create, rule like following line
	Toy Toy `gorm:"polymorphicId:OwnerID;polymorphicType:OwnerType;polymorphicValue:pigs"`
}

type Dog struct {
	ID   uint
	Name string
	Toy  Toy `gorm:"polymorphic:Owner"`
}

type Cat struct {
	ID   uint
	Name string
	Toy  Toy `gorm:"polymorphic:Owner"`
}

type Toy struct {
	ID        uint
	Name      string
	OwnerID   uint
	OwnerType string
}

func Test(db *gorm.DB) {
	tbls := []string{"cats", "dogs", "toys"}
	m := db.Migrator()
	for _, tbl := range tbls {
		if m.HasTable(tbl) {
			m.DropTable(tbl)
		}
	}

	m.AutoMigrate(&Dog{}, &Cat{}, &Toy{})

	db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})
	db.Create(&Cat{Name: "cat1", Toy: Toy{Name: "toy2"}})
}
