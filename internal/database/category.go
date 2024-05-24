package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db   *sql.DB
	ID   string
	Name string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) CreateCategory(name string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name) VALUES ($1, $2)",
		id, name)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name}, nil
}
