package tags

import (
	"fmt"

	"gorm.io/gorm"
)

// don't use tag
/*
type Dog struct {
	ID   uint
	NickName string
	Toys []Toy
}

type Toy struct {
	ID    uint
	Name  string
	DogID uint
}
*/

// use tag
/*
type Dog struct {
	ID   uint
	NickName string
	// specific toy's a column to be foreign key
	// foreign key will save Dog filed ID (it's prime key)
	Toys []Toy `gorm:"foreignKey:DogID"`
}

type Toy struct {
	ID    uint
	NickName  string
	DogID uint
}
*/

/*
type Dog struct {
	ID       uint
	NickName string
	// references tag allow you don't use prime key to be foreign key
	Toys []Toy `gorm:"foreignKey:DogNickName;references:NickName"`
}

type Toy struct {
	ID          uint
	Name        string
	DogNickName string
}
*/

/*
func Test(db *gorm.DB) {
	m := db.Migrator()
	tbls := []string{"dogs", "toys"}
	for _, tbl := range tbls {
		if m.HasTable(tbl) {
			m.DropTable(tbl)
		}
	}
	m.AutoMigrate(&Dog{}, &Toy{})

	t1 := Toy{Name: "toy1"}
	t2 := Toy{Name: "toy2"}
	d1 := Dog{NickName: "dog1", Toys: []Toy{t1, t2}}
	db.Create(&d1)
}
*/

// many to many
/*
type Dog struct {
	ID       uint
	NickName string
	// many to many tag structure has differ
	Toys []Toy `gorm:"many2many:dog_toys;foreignKey:NickName;references:Name"`
}

type Toy struct {
	ID   uint
	Name string
	// this is opposite of before
	Dogs []Dog `gorm:"many2many:toy_dogs;foreignKey:Name;references:NickName"`
}
// dog to toy: dog_toys; toy to dog: toy_dogs
*/

type Dog struct {
	ID       uint
	NickName string
	// many to many tag structure has differ
	// joinForeignKey is rename the column of join table
	// joinReferences is rename the column of join table
	Toys []Toy `gorm:"many2many:dog_toys;foreignKey:NickName;references:Name;joinForeignKey:NickName;joinReferences:Name"`
}

type Toy struct {
	ID   uint
	Name string
	// this is opposite of before, table name must be one
	Dogs []Dog `gorm:"many2many:toy_dogs;foreignKey:Name;references:NickName;joinForeignKey:Name;joinReferences:NickName"`
}

func Test2(db *gorm.DB) {
	m := db.Migrator()
	tbls := []string{"dogs", "toys", "dog_toys", "toy_dogs"}
	for _, tbl := range tbls {
		if m.HasTable(tbl) {
			m.DropTable(tbl)
		}
	}
	m.AutoMigrate(&Dog{}, &Toy{})

	t1 := Toy{Name: "toy1"}
	t2 := Toy{Name: "toy2"}
	ts := []Toy{t1, t2}

	d1 := Dog{NickName: "dog1", Toys: ts}
	d2 := Dog{NickName: "dog2", Toys: ts}

	db.Create(&d1)
	db.Create(&d2)

	ds := []Dog{d1, d2}
	t3 := Toy{Name: "toy3", Dogs: ds}

	db.Create(&t3)
}

func Test3(db *gorm.DB) {
	var dogs []Dog
	var toys []Toy
	dogs = dogs
	toys = toys

	// db.Preload("Toy").Find(&dogs)
	// fmt.Println(dogs)

	db.Preload("Dogs").Find(&toys)
	fmt.Println(toys)
}
