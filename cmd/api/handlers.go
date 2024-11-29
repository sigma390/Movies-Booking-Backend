package main

import (
	"fmt"

	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) { //accept 2 things , 1 Writer (sending what to client) and 2 Req which is type of Pointer to httpRequest
	fmt.Fprint(w, "Hello World")
}

// func (app *application) Home( w http.ResponseWriter, r*http.Request) { 
// 	fmt.Fprintf(w,"Hellow from %s", app.Domain );
// }

func (app *application) Home( w http.ResponseWriter, r *http.Request){

	//var Payload 
	var payload = struct{
		Status string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`

	}{
		Status :"active",
		Message: "Go movie Up And Running",
		Version: "1.0.0",
	}
	//convert Above to Json
	// out,err:= json.Marshal(payload);
	// if err !=nil{
	// 	fmt.Println(err);
	// }

	//headers
	// w.Header().Set("Content-Type","application/json")
	// w.WriteHeader(http.StatusOK) // Appends Previos Header Actually 
	// w.Write(out);

	_ = app.writeJSON(w, http.StatusOK , payload)

}

func (app *application) AllMovies( w http.ResponseWriter, r *http.Request){
	//get All Movies
	 //array Of Movies

	 //====================> OLD MOvies <==========================
	//  movies := [] models.Movie{
	// 	{
	// 		ID:           1,
	// 		Title:        "Inception",
	// 		ReleaseDate: "2010-07-16",
	// 		RunTime:      148,
	// 		MPAARating:   "PG-13",
	// 		Description:  "A skilled thief, the absolute best in the dangerous art of extraction, steals valuable secrets from deep within the subconscious during the dream state.",
	// 		Image:        "https://example.com/images/inception.jpg",
	// 		CreatedAt:    time.Now(),
	// 		UpdatedAt:    time.Now(),
	// 	},
	// 	{
	// 		ID:           2,
	// 		Title:        "The Dark Knight",
	// 		ReleaseDate: "2008-07-18",
	// 		RunTime:      152,
	// 		MPAARating:   "PG-13",
	// 		Description:  "Batman raises the stakes in his war on crime, setting out to dismantle the remaining criminal organizations that plague Gotham City.",
	// 		Image:        "https://example.com/images/dark-knight.jpg",
	// 		CreatedAt:    time.Now(),
	// 		UpdatedAt:    time.Now(),
	// 	},
	// }
	

	//========================> NEW Movies <===============================
	
	movies,err := app.DB.AllMovies();
	if err!=nil{
       app.errorJSON(w,err)
		return 
    }
	//===========> Old Method <==================
	//convert to json
	// out,err := json.Marshal(movies);
	// if err!=nil{
	// 	fmt.Println(err);
	// }
	//Headers
	// w.Header().Set("Content-Type","application/json")
	// w.WriteHeader(http.StatusOK)

	//send to client
	// w.Write(out);

	//================> New Methos <================
	_ = app.writeJSON(w,http.StatusOK,movies)
}