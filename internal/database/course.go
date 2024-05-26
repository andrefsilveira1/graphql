package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db         *sql.DB
	ID         string
	Name       string
	CategoryID string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) CreateCourse(name, categoryID string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec(`INSERT INTO courses (id, name, category_id) VALUS ($1,$2,$3)`, id, name, categoryID)
	if err != nil {
		return nil, err
	}

	return &Course{
		ID:         id,
		Name:       name,
		CategoryID: categoryID,
	}, nil
}

func (c *Course) ListCourses() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id, name, category_id string
		if err := rows.Scan(&id, &name, &category_id); err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Name: name, CategoryID: category_id})
	}

	return courses, nil
}
