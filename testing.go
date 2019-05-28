package main

import (
    _ "github.com/ibmdb/go_ibm_db"
    "database/sql"
    "fmt"
)

func Create_Con(con string) *sql.DB{
 db, err:=sql.Open("go_ibm_db", con)
    if err != nil{
        
		fmt.Println(err)
		return nil
	}
	return db
}

//creating a table

func create(db *sql.DB) error{
    _,err:=db.Exec("DROP table SAMPLE")
	if(err!=nil){
    _,err:=db.Exec("create table SAMPLE(ID int,NAME varchar(20),LOCATION float,POSITION varchar(20))")
    if err != nil{
        return err
    }
	} else {
    _,err:=db.Exec("create table SAMPLE(ID int,NAME varchar(20),LOCATION float,POSITION varchar(20))")
    if err != nil{
        return err
    }
	}
	fmt.Println("TABLE CREATED IN DATABASE")
    return nil
}

//inserting row

func insert(db *sql.DB) error{
    st, err:=db.Prepare("Insert into SAMPLE(ID,NAME,LOCATION,POSITION) values(3242,'mike',123.24,'manager')")
    if err != nil{
        return err
    }
    st.Query()
    return nil
}

//this api selects the data from the table and prints it

func display(db *sql.DB) error{
    st, err:=db.Prepare("select * from SAMPLE")
    if err !=nil{
        return err
    }
    err=execquery(st)
    if err!=nil{
        return err
    }
    return nil
}


func execquery(st *sql.Stmt) error{
    rows,err :=st.Query()
    if err != nil{
        return err
    }
	cols, _ := rows.Columns()
    fmt.Printf("%s    %s   %s    %s\n",cols[0],cols[1],cols[2],cols[3])
    fmt.Println("-------------------------------------")
    defer rows.Close()
    for rows.Next(){
		var t int
		var m float64
        var x,n string
        err = rows.Scan(&t,&x,&m,&n)
        if err != nil{
            return err
        }
        fmt.Printf("%v  %v   %v         %v\n",t,x,m,n)
    }
    return nil
}

func main(){
    con := "HOSTNAME=dashdb-entry-yp-lon02-01.services.eu-gb.bluemix.net;DATABASE=BLUDB;PORT=50000;UID=dash13365;PWD=******"
	type Db *sql.DB
	var re Db
	re=Create_Con(con)
    err:=create(re)
	if err != nil{
        fmt.Println(err)
    }
    err=insert(re)
    if err != nil{
        fmt.Println(err)
    }
    err=display(re)
    if err != nil{
        fmt.Println(err)
    }
}
