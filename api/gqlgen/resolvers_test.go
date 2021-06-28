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
	user := pg.User{
		ID:    1,
		Name:  "Test",
		Email: "test@example.com",
	}
	userID := int64(1)
	repo.On("CreateList", mock.Anything, mock.AnythingOfType("pg.CreateListParams")).Return(list, nil)
	repo.On("GetUser", mock.Anything, mock.AnythingOfType("int64")).Return(user, nil)
	auth.On("GetUserID", mock.Anything).Return(userID, nil)
	var response struct {
		CreateList struct {
			ID   int64
			Name string
			User pg.User
		}
	}
	query := `
		mutation CreateList($input: ListInput!) {
			createList(data: $input) {
				id
				name
				user {
					id
					name
					email
				}
			}
		}
	`
	c.MustPost(query, &response, client.Var("input", ListInput{Name: list.Name}))
	require.Equal(t, int64(1), response.CreateList.ID)
	require.Equal(t, "Test", response.CreateList.Name)
	require.Equal(t, user, response.CreateList.User)
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

func TestGetList_Error_NoUser(t *testing.T) {
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
	user := pg.User{
		ID:    1,
		Name:  "Test",
		Email: "test@example.com",
	}
	repo.On("GetList", mock.Anything, mock.AnythingOfType("int64")).Return(list, nil)
	user_error := errors.New("no user")
	repo.On("GetUser", mock.Anything, mock.AnythingOfType("int64")).Return(user, user_error)
	var response struct {
		List *struct {
			ID   int64
			Name string
			User pg.User
		}
	}
	query := `
		query List($input: ID!) {
			list(id: $input) {
				id
				name
				user {
					id
					name
					email
				}
			}
		}
	`
	err = c.Post(query, &response, client.Var("input", int64(1)))
	require.Nil(t, response.List)
	require.Error(t, err)
}

func TestGetList_Error_NoList(t *testing.T) {
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
	db_error := sql.ErrNoRows
	repo.On("GetList", mock.Anything, mock.AnythingOfType("int64")).Return(list, db_error)
	var response struct {
		List *struct {
			ID   int64
			Name string
			User pg.User
		}
	}
	query := `
		query List($input: ID!) {
			list(id: $input) {
				id
				name
				user {
					id
					name
					email
				}
			}
		}
	`
	err = c.Post(query, &response, client.Var("input", int64(1)))
	require.Nil(t, response.List)
	require.Error(t, err)
}

func TestGetList_Success(t *testing.T) {
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
	item := pg.Item{
		ID:     1,
		Name:   "Test",
		Source: "Test",
	}
	listItem1 := pg.ListItem{
		ID:        1,
		Quantity:  sql.NullInt64{Valid: false},
		Collected: false,
		ItemID:    item.ID,
		ListID:    1,
	}
	listItem2 := pg.ListItem{
		ID:        2,
		Quantity:  sql.NullInt64{Int64: 1, Valid: true},
		Collected: false,
		ItemID:    item.ID,
		ListID:    1,
	}
	user := pg.User{
		ID:    1,
		Name:  "Test",
		Email: "test@example.com",
	}
	repo.On("GetList", mock.Anything, mock.AnythingOfType("int64")).Return(list, nil)
	repo.On("GetUser", mock.Anything, mock.AnythingOfType("int64")).Return(user, nil)
	repo.On("GetListListItems", mock.Anything, mock.AnythingOfType("int64")).Return([]pg.ListItem{listItem1, listItem2}, nil)
	repo.On("GetItem", mock.Anything, mock.AnythingOfType("int64")).Return(item, nil)
	var response struct {
		List *struct {
			ID    int64
			Name  string
			User  pg.User
			Items []*struct {
				ID        int64
				Quantity  *int64
				Collected bool
				Item      pg.Item
				List      pg.List
			}
		}
	}
	query := `
		query List($input: ID!) {
			list(id: $input) {
				id
				name
				user {
					id
					name
					email
				}
				items {
					id
					quantity
					collected
					item {
						id
						name
						source
					}
					list {
						id
						name
					}
				}
			}
		}
	`
	c.MustPost(query, &response, client.Var("input", int64(1)))
	require.Equal(t, response.List.Name, "Test")
	require.Equal(t, response.List.User.ID, int64(1))
	listItems := response.List.Items
	require.Equal(t, listItems[0].ID, int64(1))
	require.Nil(t, listItems[0].Quantity)
	require.Equal(t, *listItems[1].Quantity, int64(1))
	require.Equal(t, listItems[0].Item.ID, int64(1))
	require.Equal(t, listItems[0].List.Name, "Test")
}

func TestGetListListItems_Error_Database(t *testing.T) {
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
	db_error := errors.New("db_error")
	repo.On("GetList", mock.Anything, mock.AnythingOfType("int64")).Return(list, nil)
	repo.On("GetListListItems", mock.Anything, mock.AnythingOfType("int64")).Return([]pg.ListItem{}, db_error)
	var response struct {
		List *struct {
			ID    int64
			Name  string
			Items *[]pg.ListItem
		}
	}
	query := `
		query List($input: ID!) {
			list(id: $input) {
				id
				name
				items {
					id
					quantity
					collected
				}
			}
		}
	`
	err = c.Post(query, &response, client.Var("input", int64(1)))
	require.Nil(t, response.List.Items)
	require.Error(t, err)
}

func TestGetListListItems_NoListItems(t *testing.T) {
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
	repo.On("GetList", mock.Anything, mock.AnythingOfType("int64")).Return(list, nil)
	repo.On("GetListListItems", mock.Anything, mock.AnythingOfType("int64")).Return([]pg.ListItem{}, nil)
	var response struct {
		List *struct {
			ID    int64
			Name  string
			Items []*struct {
				ID        int64
				Quantity  *int64
				Collected bool
			}
		}
	}
	query := `
		query List($input: ID!) {
			list(id: $input) {
				id
				name
				items {
					id
					quantity
					collected
				}
			}
		}
	`
	err = c.Post(query, &response, client.Var("input", int64(1)))
	require.Empty(t, response.List.Items)
	require.Nil(t, err)
}

func TestGetListItemItem_Error_Database(t *testing.T) {
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
	item := pg.Item{
		ID:     1,
		Name:   "Test",
		Source: "Test",
	}
	listItem := pg.ListItem{
		ID:        1,
		Quantity:  sql.NullInt64{Valid: false},
		Collected: false,
		ItemID:    item.ID,
		ListID:    1,
	}
	item_error := errors.New("item error")
	repo.On("GetList", mock.Anything, mock.AnythingOfType("int64")).Return(list, nil)
	repo.On("GetListListItems", mock.Anything, mock.AnythingOfType("int64")).Return([]pg.ListItem{listItem}, nil)
	repo.On("GetItem", mock.Anything, mock.AnythingOfType("int64")).Return(item, item_error)
	var response struct {
		List *struct {
			ID    int64
			Name  string
			Items []*struct {
				ID        int64
				Quantity  *int64
				Collected bool
				Item      pg.Item
			}
		}
	}
	query := `
		query List($input: ID!) {
			list(id: $input) {
				id
				name
				items {
					id
					quantity
					collected
					item {
						id
						name
						source
					}
				}
			}
		}
	`
	err = c.Post(query, &response, client.Var("input", int64(1)))
	require.Error(t, err)
	require.Nil(t, response.List.Items[0])
}
