package customize

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Info struct {
	ID   uint
	Name string `gorm:"type:text"` // specific data type in mysql
}

type Msg struct {
	ID   uint
	Info Info
	Args Args
}

type Args []string

/*
driver, database driver
error, if throw error, Info will not be saved
*/
func (c Info) Value() (driver.Value, error) {
	// save to database by json format
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}

/*
 */
func (c *Info) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("mismatch type")
	}
	json.Unmarshal(bytes, c)
	return nil
}

func (a Args) Value() (driver.Value, error) {
	if len(a) > 0 {
		s := a[0]
		for _, v := range a[1:] {
			s += "," + v
		}
		return s, nil
	} else {
		return "", nil
	}
}

func (a *Args) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("mismatch type")
	}
	*a = strings.Split(string(bytes), ",")
	return nil
}

func Test(db *gorm.DB) {
	m := db.Migrator()
	tbls := []string{"infos", "msgs"}
	for _, v := range tbls {
		if m.HasTable(v) {
			m.DropTable(v)
		}
		// can't create, because don't know table structure
		// m.AutoMigrate(v)
	}
	m.AutoMigrate(&Info{}, &Msg{})
	// infos not have data, msgs have complete data
	db.Create(&Msg{Info: Info{Name: "info1"}, Args: Args{"a", "b", "c"}})

	var msg Msg
	db.Find(&msg)
	fmt.Println(msg)
}
