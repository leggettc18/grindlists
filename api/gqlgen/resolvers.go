package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
)

type Resolver struct{}

func (r *mutationResolver) CreateUser(ctx context.Context, data UserInput) (*User, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int64, data UserInput) (*User, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int64) (*User, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateList(ctx context.Context, data ListInput) (*List, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateList(ctx context.Context, id int64, data ListInput) (*List, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteList(ctx context.Context, id int64) (*List, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateListItem(ctx context.Context, itemData ItemInput, listItemdata ListItemInput) (*Item, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateItem(ctx context.Context, id int64, data ItemInput) (*Item, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteItem(ctx context.Context, id int64) (*Item, error) {
	panic("not implemented")
}

func (r *mutationResolver) SetListItem(ctx context.Context, data ListItemInput) (*ListItem, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateListItem(ctx context.Context, id int64, data ListItemInput) (*ListItem, error) {
	panic("not implemented")
}

func (r *mutationResolver) UnsetListItem(ctx context.Context, id int64) (*ListItem, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context, id int64) (*User, error) {
	panic("not implemented")
}

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	panic("not implemented")
}

func (r *queryResolver) List(ctx context.Context, id int64) (*List, error) {
	panic("not implemented")
}

func (r *queryResolver) Lists(ctx context.Context) ([]List, error) {
	panic("not implemented")
}

func (r *queryResolver) Item(ctx context.Context, id int64) (*Item, error) {
	panic("not implemented")
}

func (r *queryResolver) Items(ctx context.Context) ([]Item, error) {
	panic("not implemented")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
