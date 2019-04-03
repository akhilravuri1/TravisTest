package main

import (
	"database/sql"
	"fmt"
	_ "github.com/ibmdb/go_ibm_db"
)

func Create_Con(con string) *sql.DB {
	fmt.Println("create conn func")
	db, err := sql.Open("go_ibm_db", con)
	if err != nil {

		fmt.Println(err)
		return nil
	}
	return db
}

//creating a table

func create(db *sql.DB) error {
	_, err := db.Exec("DROP table SAMPLE")
	fmt.Println(err)
	fmt.Println("This is table create")
	if err != nil {
		_, err := db.Exec("create table SAMPLE(ID varchar(20),NAME varchar(20),LOCATION varchar(20),POSITION varchar(20))")
		if err != nil {
		    fmt.Println("this is inside error") 
			return err
		}
	} else {
		_, err := db.Exec("create table SAMPLE(ID varchar(20),NAME varchar(20),LOCATION varchar(20),POSITION varchar(20))")
		if err != nil {
			return err
		}
	}
	fmt.Println("TABLE CREATED")
	return nil
}

func main() {
	fmt.Println("This is main")
	con := "HOSTNAME=10.0.75.1;DATABASE=go;PORT=50000;UID=rakhil;PWD=Spranu134#"
	type Db *sql.DB
	var re Db
	re = Create_Con(con)
	err := create(re)
	if err != nil {
		fmt.Println(err)
	}
}
