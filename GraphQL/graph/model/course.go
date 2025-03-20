package model

type Course struct {
	ID          string
	Name        string
	Description *string
	Category    *Category
}
