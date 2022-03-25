package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/ahmedkhaeld/graphql-server/graph/generated"
	"github.com/ahmedkhaeld/graphql-server/graph/model"
	"github.com/ahmedkhaeld/graphql-server/repository"
	"math/rand"
	"strconv"
)

var videoRepo = repository.NewVideoRepository()

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {

	video := &model.Video{
		ID:     strconv.Itoa(rand.Int()),
		Title:  input.Title,
		URL:    input.URL,
		Author: &model.User{ID: input.UserID, Name: "user" + input.UserID},
	}
	videoRepo.Save(video)
	return video, nil
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	return videoRepo.FindAll(), nil
}

func (r *subscriptionResolver) VideoAdded(ctx context.Context, repoFullName string) (<-chan *model.Video, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
