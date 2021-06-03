package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"time"

	"github.com/leggettc18/grindlists/api/pg"
)

type Resolver struct{
	Repository pg.Repository
}

func (r *listResolver) User(ctx context.Context, obj *pg.List) (*pg.User, error) {
	panic("not implemented")
}

func (r *listResolver) Items(ctx context.Context, obj *pg.List) ([]*pg.ListItem, error) {
	panic("not implemented")
}

func (r *listItemResolver) Quantity(ctx context.Context, obj *pg.ListItem) (*int, error) {
	panic("not implemented")
}

func (r *listItemResolver) List(ctx context.Context, obj *pg.ListItem) (*pg.List, error) {
	panic("not implemented")
}

func (r *listItemResolver) Item(ctx context.Context, obj *pg.ListItem) (*pg.Item, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateUser(ctx context.Context, data UserInput) (*pg.User, error) {
	user, err := r.Repository.CreateUser(ctx, pg.CreateUserParams{
		Name: data.Name,
		Email: data.Email,
		HashedPassword: []byte(data.Password), // will replace with actual hashed password once graphql endpoint is working.
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int64, data UserInput) (*pg.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int64) (*pg.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateList(ctx context.Context, data ListInput) (*pg.List, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateList(ctx context.Context, id int64, data ListInput) (*pg.List, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteList(ctx context.Context, id int64) (*pg.List, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateListItem(ctx context.Context, itemData ItemInput, listItemdata ListItemInput) (*pg.Item, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateItem(ctx context.Context, id int64, data ItemInput) (*pg.Item, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteItem(ctx context.Context, id int64) (*pg.Item, error) {
	panic("not implemented")
}

func (r *mutationResolver) SetListItem(ctx context.Context, data ListItemInput) (*pg.ListItem, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateListItem(ctx context.Context, id int64, data ListItemInput) (*pg.ListItem, error) {
	panic("not implemented")
}

func (r *mutationResolver) UnsetListItem(ctx context.Context, id int64) (*pg.ListItem, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context, id int64) (*pg.User, error) {
	panic("not implemented")
}

func (r *queryResolver) Users(ctx context.Context) ([]pg.User, error) {
	return r.Repository.ListUsers(ctx)
}

func (r *queryResolver) List(ctx context.Context, id int64) (*pg.List, error) {
	panic("not implemented")
}

func (r *queryResolver) Lists(ctx context.Context) ([]pg.List, error) {
	panic("not implemented")
}

func (r *queryResolver) Item(ctx context.Context, id int64) (*pg.Item, error) {
	panic("not implemented")
}

func (r *queryResolver) Items(ctx context.Context) ([]pg.Item, error) {
	panic("not implemented")
}

func (r *userResolver) Lists(ctx context.Context, obj *pg.User) ([]*pg.List, error) {
	panic("not implemented")
}

// List returns ListResolver implementation.
func (r *Resolver) List() ListResolver { return &listResolver{r} }

// ListItem returns ListItemResolver implementation.
func (r *Resolver) ListItem() ListItemResolver { return &listItemResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type listResolver struct{ *Resolver }
type listItemResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }