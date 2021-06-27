package gqlgen

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/leggettc18/grindlists/api/app"
	authMocks "github.com/leggettc18/grindlists/api/mocks/auth"
	repoMocks "github.com/leggettc18/grindlists/api/mocks/pg"
	"github.com/leggettc18/grindlists/api/pg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateList_Success(t *testing.T) {
	auth := new(authMocks.AuthService)
	repo := new(repoMocks.Repository)
	app, err := app.New()
	assert.Nil(t, err)
	c := client.New(NewHandler(repo, *app, auth))
	list := pg.List{
		ID:        1,
		Name:      "Test",
		UserID:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: sql.NullTime{Valid: false},
	}
	userID := int64(1)
	repo.On("CreateList", mock.Anything, mock.AnythingOfType("pg.CreateListParams")).Return(list, nil)
	auth.On("GetUserID", mock.Anything).Return(userID, nil)
	var response struct {
		CreateList struct {
			ID   int64
			Name string
		}
	}
	query := `
		mutation CreateList($input: ListInput!) {
			createList(data: $input) {
				id
				name
			}
		}
	`
	c.MustPost(query, &response, client.Var("input", ListInput{Name: list.Name}))
	require.Equal(t, int64(1), response.CreateList.ID)
	require.Equal(t, "Test", response.CreateList.Name)
}
func TestCreateList_Error_NotAuthenticated(t *testing.T) {
	auth := new(authMocks.AuthService)
	repo := new(repoMocks.Repository)
	app, err := app.New()
	assert.Nil(t, err)
	c := client.New(NewHandler(repo, *app, auth))
	list := pg.List{
		ID:        1,
		Name:      "Test",
		UserID:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: sql.NullTime{Valid: false},
	}
	userID := int64(-1)
	auth.On("GetUserID", mock.Anything).Return(userID, errors.New("not authenticated"))
	repo.On("CreateList", mock.Anything, mock.AnythingOfType("pg.CreateListParams")).Return(list, nil)
	var response struct {
		CreateList *struct {
			ID   int64
			Name string
		}
	}
	query := `
		mutation CreateList($input: ListInput!) {
			createList(data: $input) {
				id
				name
			}
		}
	`
	err = c.Post(query, &response, client.Var("input", ListInput{Name: list.Name}))
	require.Nil(t, response.CreateList)
	require.EqualError(t, err, "[{\"message\":\"not authenticated\",\"path\":[\"createList\"]}]")
}

func TestCreateList_Error_Database(t *testing.T) {
	auth := new(authMocks.AuthService)
	repo := new(repoMocks.Repository)
	app, err := app.New()
	assert.Nil(t, err)
	c := client.New(NewHandler(repo, *app, auth))
	list := pg.List{
		ID:        1,
		Name:      "Test",
		UserID:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: sql.NullTime{Valid: false},
	}
	userID := int64(1)
	db_error := errors.New("db_error")
	auth.On("GetUserID", mock.Anything).Return(userID, nil)
	repo.On("CreateList", mock.Anything, mock.AnythingOfType("pg.CreateListParams")).Return(list, db_error)
	var response struct {
		CreateList *struct {
			ID   int64
			Name string
		}
	}
	query := `
		mutation CreateList($input: ListInput!) {
			createList(data: $input) {
				id
				name
			}
		}
	`
	err = c.Post(query, &response, client.Var("input", ListInput{Name: list.Name}))
	require.Nil(t, response.CreateList)
	require.Error(t, err)
}
