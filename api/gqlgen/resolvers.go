package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"errors"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/leggettc18/grindlists/api/app"
	"github.com/leggettc18/grindlists/api/auth"
	"github.com/leggettc18/grindlists/api/pg"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Resolver struct {
	Repository pg.Repository
	App        app.App
	Auth       auth.AuthService
}

func (r *listResolver) User(ctx context.Context, obj *pg.List) (*pg.User, error) {
	user, err := r.Repository.GetUser(ctx, obj.UserID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *listResolver) Items(ctx context.Context, obj *pg.List) ([]pg.ListItem, error) {
	listItems, err := r.Repository.GetListListItems(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return listItems, nil
}

func (r *listItemResolver) Quantity(ctx context.Context, obj *pg.ListItem) (*int, error) {
	var quantity int
	if obj.Quantity.Valid {
		quantity = int(obj.Quantity.Int64)
		return &quantity, nil
	}
	return nil, nil
}

func (r *listItemResolver) Item(ctx context.Context, obj *pg.ListItem) (*pg.Item, error) {
	item, err := r.Repository.GetItem(ctx, obj.ItemID)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *queryResolver) Me(ctx context.Context) (*pg.User, error) {
	userID, err := r.Auth.GetUserID(ctx)
	if err != nil {
		return nil, errors.New("not authenticated")
	}
	user, err := r.Repository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mutationResolver) Login(ctx context.Context, data LoginInput) (*pg.User, error) {
	user, err := r.Repository.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}
	valid, err := auth.VerifyPasswordHash(data.Password, user.HashedPassword)
	if err != nil {
		return nil, err
	}
	if !valid {
		graphql.AddError(ctx, &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "Incorrect Password.",
			Extensions: map[string]interface{}{
				"field": "password",
			},
		})
		return nil, nil
	}
	token, err := auth.CreateToken(user.ID, r.App.Config.SecretKey)
	if err != nil {
		return nil, err
	}
	saveErr := auth.CacheAuth(user.ID, token)
	if saveErr != nil {
		return nil, saveErr
	}
	cookieAccess := auth.GetCookieAccess(ctx)
	cookieAccess.SetToken("jwtAccess", token.AccessToken, time.Unix(token.AtExpires, 0))
	cookieAccess.SetToken("jwtRefresh", token.RefreshToken, time.Unix(token.RtExpires, 0))
	return &user, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (*LogoutOutput, error) {
	userID, ok := ctx.Value(auth.UserIDKey).(int64)
	if !ok {
		return &LogoutOutput{Succeeded: false}, errors.New("not authenticated")
	}
	accessUuid, ok := ctx.Value(auth.AccessUuidKey).(string)
	if !ok {
		return &LogoutOutput{Succeeded: false}, errors.New("access uuid not present in context")
	}
	deleted, err := auth.DeleteAuth("access_token", accessUuid)
	if err != nil || deleted == 0 {
		return &LogoutOutput{Succeeded: false}, errors.New("not authenticated")
	}
	refreshUuid, ok := ctx.Value(auth.RefreshUuidKey).(string)
	if !ok {
		return &LogoutOutput{
			Succeeded: false,
		}, errors.New("refresh uuid not present in context")
	}
	deleted, err = auth.DeleteAuth("refresh_token", refreshUuid)
	if err != nil || deleted == 0 {
		return &LogoutOutput{Succeeded: false}, errors.New("not authenticated")
	}
	cookieAccess := auth.GetCookieAccess(ctx)
	cookieAccess.RemoveToken("jwtAccess")
	cookieAccess.RemoveToken("jwtRefresh")
	return &LogoutOutput{UserID: userID, Succeeded: true}, nil
}

func (r *mutationResolver) Register(ctx context.Context, data UserInput) (*pg.User, error) {
	hashedPassword, err := auth.GetPasswordHash(data.Password)
	if err != nil {
		return nil, err
	}
	user, err := r.Repository.CreateUser(ctx, pg.CreateUserParams{
		Name:           data.Name,
		Email:          data.Email,
		HashedPassword: hashedPassword,
		CreatedAt:      time.Now(),
	})
	if err != nil {
		return nil, err
	}
	token, err := auth.CreateToken(user.ID, r.App.Config.SecretKey)
	if err != nil {
		return nil, err
	}
	saveErr := auth.CacheAuth(user.ID, token)
	if saveErr != nil {
		return nil, saveErr
	}
	cookieAccess := auth.GetCookieAccess(ctx)
	cookieAccess.SetToken("jwtAccess", token.AccessToken, time.Unix(token.AtExpires, 0))
	cookieAccess.SetToken("jwtRefresh", token.RefreshToken, time.Unix(token.RtExpires, 0))
	return &user, nil
}

func (r *mutationResolver) Refresh(ctx context.Context) (*pg.User, error) {
	userID, ok := ctx.Value(auth.RefreshUserIDKey).(int64)
	if !ok {
		return nil, errors.New("not authenticated (no user id in context)")
	}
	user, err := r.Repository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	refreshUuid, ok := ctx.Value(auth.RefreshUuidKey).(string)
	if !ok {
		return nil, errors.New("refresh token uuid not present in context")
	}
	deleted, err := auth.DeleteAuth("refresh_token", refreshUuid)
	if err != nil || deleted == 0 {
		return nil, errors.New("not authenticated (no refresh token in cache)")
	}
	token, err := auth.CreateToken(user.ID, r.App.Config.SecretKey)
	if err != nil {
		return nil, err
	}
	saveErr := auth.CacheAuth(user.ID, token)
	if saveErr != nil {
		return nil, saveErr
	}
	cookieAccess := auth.GetCookieAccess(ctx)
	cookieAccess.SetToken("jwtAccess", token.AccessToken, time.Unix(token.AtExpires, 0))
	cookieAccess.SetToken("jwtRefresh", token.RefreshToken, time.Unix(token.RtExpires, 0))
	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int64, data UserInput) (*pg.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int64) (*pg.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateList(ctx context.Context, data ListInput) (*pg.List, error) {
	userID, err := r.Auth.GetUserID(ctx)
	if err != nil {
		return nil, errors.New("not authenticated")
	}
	list, err := r.Repository.CreateList(ctx, pg.CreateListParams{
		Name:      data.Name,
		UserID:    userID,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *mutationResolver) UpdateList(ctx context.Context, id int64, data ListInput) (*pg.List, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteList(ctx context.Context, id int64) (*pg.List, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateListItem(ctx context.Context, listItemData CreateListItemInput) (*pg.Item, error) {
	item, err := r.Repository.CreateListItem(ctx, pg.CreateItemParams{
		Name:   listItemData.Name,
		Source: listItemData.Source,
	}, pg.SetListItemParams{
		Quantity:  pg.IntPtrToNullInt64(listItemData.Quantity),
		Collected: listItemData.Collected,
		ListID:    listItemData.ListID,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return item, nil
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

func (r *mutationResolver) Heart(ctx context.Context, list_id int64) (*pg.List, error) {
	user_id, err := r.Auth.GetUserID(ctx)
	if err != nil {
		return nil, errors.New("not authenticated")
	}
	err = r.Repository.SetListHeart(ctx, pg.SetListHeartParams{
		ListID:    list_id,
		UserID:    user_id,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	list, err := r.Repository.GetList(ctx, list_id)
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *queryResolver) User(ctx context.Context, id int64) (*pg.User, error) {
	user, err := r.Repository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]pg.User, error) {
	return r.Repository.ListUsers(ctx)
}

func (r *queryResolver) List(ctx context.Context, id int64) (*pg.List, error) {
	list, err := r.Repository.GetList(ctx, id)
	if err != nil {
		return nil, err
	}
	return &list, err
}

func (r *queryResolver) Lists(ctx context.Context) ([]pg.List, error) {
	return r.Repository.ListLists(ctx)
}

func (r *queryResolver) Item(ctx context.Context, id int64) (*pg.Item, error) {
	item, err := r.Repository.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]pg.Item, error) {
	return r.Repository.ListItems(ctx)
}

func (r *userResolver) Lists(ctx context.Context, obj *pg.User) ([]pg.List, error) {
	userLists, err := r.Repository.GetUserLists(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return userLists, nil
}

func (r *listResolver) Hearts(ctx context.Context, obj *pg.List) (*ListHeartAggregate, error) {
	preloads := getPreloads(ctx)
	listHearts := []pg.ListHeart{}
	var err error = nil
	if contains(preloads, "hearts") {
		listHearts, err = r.Repository.GetListHearts(ctx, obj.ID)
		if err != nil {
			return nil, err
		}
	}
	var countHearts int64 = 0
	if contains(preloads, "count") {
		countHearts, err = r.Repository.CountListHearts(ctx, obj.ID)
		if err != nil {
			return nil, err
		}
	}
	var heartedByCurrentUser = false
	if contains(preloads, "by_current_user") {
		user_id, err := r.Auth.GetUserID(ctx)
		if err != nil {
			heartedByCurrentUser = false
		}
		for _, value := range listHearts {
			if value.UserID == user_id {
				heartedByCurrentUser = true
				break
			}
		}
	}
	heartAggregate := ListHeartAggregate{
		Count:         int(countHearts),
		ByCurrentUser: heartedByCurrentUser,
		Hearts:        listHearts,
	}
	return &heartAggregate, nil
}

func (r *listHeartResolver) List(ctx context.Context, obj *pg.ListHeart) (*pg.List, error) {
	list, err := r.Repository.GetList(ctx, obj.ListID)
	if err != nil {
		return nil, err
	}
	return &list, err
}
func (r *listHeartResolver) User(ctx context.Context, obj *pg.ListHeart) (*pg.User, error) {
	user, err := r.Repository.GetUser(ctx, obj.UserID)
	if err != nil {
		return nil, err
	}
	return &user, err
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

func (r *Resolver) ListHeart() ListHeartResolver { return &listHeartResolver{r} }

type listResolver struct{ *Resolver }
type listItemResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type listHeartResolver struct{ *Resolver }
