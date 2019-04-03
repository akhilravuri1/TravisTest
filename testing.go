package main

import (
	_ "github.com/ibmdb/go_ibm_db"
    "database/sql"
    "fmt"
)

func main(){
    con:="HOSTNAME=host;DATABASE=name;PORT=number;UID=username;PWD=password"
	db, err:=sql.Open("go_ibm_db", con)
    if err != nil{  
		fmt.Println(err)
	}
	db.Close()
}