package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) CreateCourse(name, description string) (*Course, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO courses (id, name, description) VALUES ($1, $2, $3)", id, name, description)
	if err != nil {
		return &Course{}, err
	}

	return &Course{ID: id, Name: name, Description: description}, nil
}
