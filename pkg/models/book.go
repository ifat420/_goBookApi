package models

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"Author"`
	Desc   string `json:"Desc"`
}
