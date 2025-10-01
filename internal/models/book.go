package models

import "time"

type Book struct {
	ID          uint   `json:"id" db:"id"`
	Title       string `json: "title" db:"title"`
	Author      string `json: "author" db:"author"`
	Description string `json:"description" db:"description"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
