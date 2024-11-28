package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
	DSN string //data Source Name
	//=> Old <==
	// DB *sql.DB //Pool of databases
	DB repository.DatabaseRepo
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
	//Old WAY
	// app.DB = conn; //this is Basically Returning A pool of database
	//Before Existing We Must Close Connection else , its A respource Leak
	// defer app.DB.Close(); //defer means Execute All under this , but before  This line Content Differ

	//new Way 
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	app.Domain = "example.com"
	log.Println("Starting Apllication on Port ",port);
	// http.HandleFunc("/", Hello); old way Default mux 
	//start a Web Server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port),app.routes())
	if err != nil{
		log.Fatal(err);
	}
}