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

func (c *Category) ListCategories() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []Category{}
	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name})
	}

	return categories, nil
}

func (c *Category) FindByCourseId(courseId string) (Category, error) {
	var id, name string
	err := c.db.QueryRow("SELECT c.id, c.name, FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = $1", courseId).Scan(&id, &name)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name}, nil
}
