package models

type Order struct {
	CommonModel
	Name        string
	Description string
	Quantity    uint
}
