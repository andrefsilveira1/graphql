package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"

	"github.com/andrefsilveira1/graphql/graph/model"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category, err := r.CategoryDB.CreateCategory(input.Name)
	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.Newcourse) (*model.Course, error) {
	course, err := r.CourseDB.CreateCourse(input.Name, *input.Description, input.CategoryID)
	if err != nil {
		return nil, err
	}
	return &model.Course{
		ID:          course.ID,
		Name:        course.Name,
		Description: &course.Description,
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	categories, err := r.CategoryDB.ListCategories()
	if err != nil {
		return nil, err
	}

	var categoriesModel []*model.Category

	for _, category := range categories {
		categoriesModel = append(categoriesModel, &model.Category{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return categoriesModel, nil
}

// Course is the resolver for the course field.
func (r *queryResolver) Course(ctx context.Context) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented: Course - course"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
