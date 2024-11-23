package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// set Application Config (db connection)
	var app application
	//read From Command Line (flags)

	//connect to DB
	app.Domain = "example.com"
	log.Println("Starting Apllication on Port ",port);
	// http.HandleFunc("/", Hello); old way Default mux 
	//start a Web Server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port),app.routes())
	if err != nil{
		log.Fatal(err);
	}
}