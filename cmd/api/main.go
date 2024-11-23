package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
	DSN string //data Source Name
	DB *sql.DB
}

func main() {
	// set Application Config (db connection)
	var app application
	//read From Command Line (flags)
	flag.StringVar(&app.DSN,"dsn","host=localhost port=5432 user=postgres password=postgres dbname = movies sslmode = disable timezone=UTC connect_timeout=5","postgres connection string")
	flag.Parse();

	//connect to DB
	conn,err := app.connectToDB()
	if err!=nil{
        log.Fatal(err);
    }
	app.DB = conn;


	app.Domain = "example.com"
	log.Println("Starting Apllication on Port ",port);
	// http.HandleFunc("/", Hello); old way Default mux 
	//start a Web Server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port),app.routes())
	if err != nil{
		log.Fatal(err);
	}
}