package main

import (
	"fmt"
	"log"
)

func getOne(id int) (data, error) {
	d := data{}
	// statement := fmt.Sprintf("SELECT Id, Name, Email FROM testdb.users where Id=%d", id)
	// err := db.QueryRow(statement).Scan(&a.ID, &a.name, &a.email)
	err := db.QueryRow(`SELECT Id, Name, Email FROM testdb.users where Id=?`,
		id).Scan(&d.ID, &d.name, &d.email)

	return d, err
}

func getMany(id int) ([]data, error) {
	statement := fmt.Sprintf("SELECT Id, Name, Email FROM testdb.users where Id=%d", id)
	rows, err := db.Query(statement)
	var ds []data
	for rows.Next() {
		d := data{}
		err = rows.Scan(&d.ID, &d.name, &d.email)
		if err != nil {
			log.Fatalln(err.Error())
		}
		ds = append(ds, d)
	}
	return ds, err
}

func (d *data) Update() error {
	statement := fmt.Sprintf("UPDATE testdb.users SET Name='%s' WHERE Id='%d'",
		d.name, d.ID)
	_, err := db.Exec(statement)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return err
}

func (d *data) Delete() error {
	statement := fmt.Sprintf("DELETE FROM testdb.users WHERE Id='%d'",
		d.ID)
	_, err := db.Exec(statement)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return err
}

func (d *data) Insert() error {
	statement :=
		`INSERT INTO testdb.users (Name, Email) VALUES (?, ?);`
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer stmt.Close()

	rs, err := stmt.Exec(d.name, d.email)

	if err != nil {
		fmt.Println("insert error")
		log.Fatalln(err.Error())
	}

	lastInsertID, err := rs.LastInsertId()
	if err != nil {
		fmt.Println("last insert id error")
		log.Fatalln(err.Error())
	}
	d.ID = int(lastInsertID)
	fmt.Println("last insert id:", lastInsertID)

	return err
}
