// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgen

import (
	"github.com/leggettc18/grindlists/api/pg"
)

type CreateListItemInput struct {
	Name      string `json:"name"`
	Source    string `json:"source"`
	Quantity  *int   `json:"quantity"`
	Collected bool   `json:"collected"`
	ListID    int64  `json:"list_id"`
}

type ItemInput struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

type ListHeartAggregate struct {
	Count         int            `json:"count"`
	ByCurrentUser bool           `json:"by_current_user"`
	Hearts        []pg.ListHeart `json:"hearts"`
}

type ListInput struct {
	Name string `json:"name"`
}

type ListItemInput struct {
	Quantity  *int  `json:"quantity"`
	Collected bool  `json:"collected"`
	ListID    int64 `json:"list_id"`
	ItemID    int64 `json:"item_id"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogoutOutput struct {
	UserID    int64 `json:"user_id"`
	Succeeded bool  `json:"succeeded"`
}

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
