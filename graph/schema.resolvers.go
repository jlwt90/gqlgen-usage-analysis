package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"
	"fmt"

	"github.com/jlwt90/gqlgen-usage-analysis/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID:   "T1",
			Text: "t1",
			Done: false,
			User: nil,
		},
		{
			ID:   "T2",
			Text: "t2",
			Done: true,
			User: nil,
		},
		{
			ID:   "T3",
			Text: "t3",
			Done: true,
			User: nil,
		},
	}, nil
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	var obj *model.Todo
	if id == "T1" {
		obj = &model.Todo{
			ID:   "T1",
			Text: "t1",
			Done: false,
			User: nil,
		}
	} else if id == "T2" {
		obj = &model.Todo{
			ID:   "T2",
			Text: "t2",
			Done: true,
			User: nil,
		}
	} else if id == "T3" {
		obj = &model.Todo{
			ID:   "T3",
			Text: "t3",
			Done: true,
			User: nil,
		}
	}
	return obj, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	if obj.ID == "T1" {
		return &model.User{
			ID:   "U1",
			Name: "u1",
		}, nil
	} else if obj.ID == "T2" {
		return &model.User{
			ID:   "U2",
			Name: "u2",
		}, nil
	}
	return &model.User{
		ID:   "U3",
		Name: "u3",
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
