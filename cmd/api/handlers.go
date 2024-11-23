package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) { //accept 2 things , 1 Writer (sending what to client) and 2 Req which is type of Pointer to httpRequest
	fmt.Fprint(w, "Hello World")
}
func (app *application) Home( w http.ResponseWriter, r*http.Request) { 
	fmt.Fprintf(w,"Hellow from %s", app.Domain );
}