package graphql_test5

import (
	"context"
	"training/graphql_test5/models"
	"training/graphql_test5/postgres"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	Database *postgres.Db
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, email string) (*models.User, error) {
	returnValue, err := r.Database.CreateUser(email)
	if err != nil {
		return nil, err
	}
	return returnValue, nil
}
func (r *mutationResolver) RemoveUser(ctx context.Context, id string) (*bool, error) {
	returnValue, err := r.Database.RemoveUser(id)
	if err != nil {
		return nil, err
	}
	return &returnValue, nil
}
func (r *mutationResolver) Follow(ctx context.Context, follower string, folowee string) (*bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) Unfollow(ctx context.Context, follower string, folowee string) (*bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePost(ctx context.Context, user string, title string, body string) (*models.Post, error) {
	returnValue, err := r.Database.CreatePost(user, title, body)
	if err != nil {
		return nil, err
	}
	return returnValue, nil
}
func (r *mutationResolver) RemovePost(ctx context.Context, id string) (*bool, error) {
	returnValue, err := r.Database.RemovePost(id)
	if err != nil {
		return nil, err
	}
	return &returnValue, nil
}
func (r *mutationResolver) CreateComment(ctx context.Context, user string, post string, title string, body string) (*models.Comment, error) {
	panic("not implemented")
}
func (r *mutationResolver) RemoveComment(ctx context.Context, id string) (*bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	returnValue, err := r.Database.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return returnValue, nil
}
func (r *queryResolver) Posts(ctx context.Context, id string) ([]models.Post, error) {
	returnValue, err := r.Database.GetPostsByUserId(id)
	if err != nil {
		return nil, err
	}

	return returnValue, nil
}
