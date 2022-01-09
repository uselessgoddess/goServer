package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"doublets-gql/graph/generated"
	"doublets-gql/graph/model"
	"strconv"
)

func (r *mutationResolver) Big(ctx context.Context, id string, data model.BigInput) ([]string, error) {
	r.Lock()
	defer r.Unlock()

	after := r.Db[id]
	r.Db[id] = data.Data
	return after, nil
}

func (r *mutationResolver) Small(ctx context.Context, id string, data model.SmallInput) ([]string, error) {
	r.Lock()
	defer r.Unlock()

	after := r.Db[id]
	r.Db[id] = []string{strconv.Itoa(data.Data)}
	return after, nil
}

func (r *queryResolver) Big(ctx context.Context, id string) (*model.Big, error) {
	r.RLock()
	defer r.RUnlock()

	if data := r.Db[id]; data != nil {
		return &model.Big{Data: data}, nil
	} else {
		return nil, nil
	}
}

func (r *queryResolver) Small(ctx context.Context, id string) (*model.Small, error) {
	r.RLock()
	defer r.RUnlock()

	if data := r.Db[id]; data != nil {
		if len(data) == 1 {
			println(data[0])
			if num, err := strconv.Atoi(data[0]); err == nil {
				return &model.Small{Data: num}, nil
			} else {
				return nil, err
			}
		} else {
			return nil, nil
		}
	} else {
		return nil, nil
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
