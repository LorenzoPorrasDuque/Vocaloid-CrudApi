package models

type Book struct {
	Title  string `json:"title" gorm:"primary_key" `
	Author string `json:"author"`
}
