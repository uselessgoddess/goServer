package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"doublets-gql/graph/generated"
	"doublets-gql/graph/model"
)

func (r *mutationResolver) InsertLinks(ctx context.Context, objects []*model.InputLink) ([]*model.Link, error) {
	var links []*model.Link
	for i := 0; i < len(objects); i++ {
		links = append(links, &model.Link{
			ID:     i+1,
			FromID: objects[i].FromID,
			ToID:   objects[i].ToID,
		})
	}

	return links, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var links []*model.Link
	for i := 0; i < 100; i++ {
		links = append(links, &model.Link{
			ID:     i+1,
			FromID: i+1,
			ToID:   i+1,
		})
	}
	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
