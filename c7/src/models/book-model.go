package models

import (
	"time"
)

type Title = string
type Author = string

// Book struct to describe book object.
// Aca declaro el modelo libro con sus tipos
// La aclaracion json me permite aclarar el nombre de la variable al ser json
type Book struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Title     Title     `json:"title" bson:"title"`
	Author    Author    `json:"author" bson:"author"`
	// ID        string    `json:"id" bson:"_id"`
}
