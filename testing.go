package main

import (
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