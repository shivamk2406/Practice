package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/shivamk2406/Practice/graphql/graph/generated"
	"github.com/shivamk2406/Practice/graphql/graph/model"
	"github.com/shivamk2406/Practice/internal/service/user"
)

// GetTenant is the resolver for the GetTenant field.
func (r *queryResolver) GetTenant(ctx context.Context, id string) (*model.Tenant, error) {
	user, err := r.repo.GetUserSubScription(ctx, &user.Model{
		ID: id,
	})
	if err != nil {
		log.Println(err)
	}

	b, err := json.Marshal(user.Data)
	if err != nil {
		log.Println(err)
	}

	var res model.Tenant

	err = json.Unmarshal(b, &res)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(res)
	return &res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
