package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/shivamk2406/Practice/graphql/graph/generated"
	"github.com/shivamk2406/Practice/graphql/graph/model"
	"github.com/shivamk2406/Practice/internal/service/user"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := r.CreateUserSubScription(ctx, &user.Model{
		ID:           uuid.New().String(),
		Name:         input.Text,
		Subscription: input.Subs,
	})
	if err != nil {
		return &model.User{}, err
	}
	return &model.User{
		ID:           user.ID,
		Name:         user.Name,
		Subscription: user.Subscription,
	}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := r.GetUserSubScription(ctx, &user.Model{
		ID: id,
	})
	if err != nil {
		return &model.User{}, err
	}

	return &model.User{
		ID:           user.ID,
		Name:         user.Name,
		Subscription: user.Subscription,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
