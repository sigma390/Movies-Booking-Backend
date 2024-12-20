package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)


func openDB(dsn string)(*sql.DB,error){ //sql.DB is a pointer pool of Database Connections
	db,err := sql.Open("pgx",dsn);
	if err!=nil{
		return nil,err;
	}
	err = db.Ping(); //tell if You are Connected Or Not to A Database
	if err!=nil{
		return nil,err;
	}
	return db,nil;
}

//function to connnect to DB basically 

func (app *application) connectToDB()(*sql.DB,error){
	connection,err:= openDB(app.DSN);
	if err!=nil{
		return nil,err;
	}
	log.Println("Connected To Postgres !!")
	return connection,nil;
}