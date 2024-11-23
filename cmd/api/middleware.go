package main

import "net/http"

//basically we Are Modifying A request That Comes in


func (app *application) enableCors (h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*");

		//Allowing All Origins
       

        if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH,OPTIONS") //Allowing All Methods
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization,Accept,X-CSRF-Token") //Allowing All Headers
			w.Header().Set("Access-Control-Allow-Credentials", "true") //Allowing Credentials

            return //Allow preflight request
        }else{
			h.ServeHTTP(w, r)

		}

        
    })
}