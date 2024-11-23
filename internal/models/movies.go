package models

import "time"
type Movie struct {
	ID           int       `json:"id`
	Title        string    `json:"title"`
	RealeaseDate string    `json:"realease_date"`
	Runtime      int       `json:"runtime"`
	MPAARating   string    `json:"mpaa_rating"`
	Description  string    `json:"description"`
	Image        string    `json:"image"`
	CreatedAt    time.Time `json:"-"` //that means Do not include in json response
	UpdatedAt    time.Time `json:"-"` //that means Do not include in json response
}


