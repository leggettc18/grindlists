package gqlgen

import (
	"database/sql"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/leggettc18/grindlists/api/app"
	authMocks "github.com/leggettc18/grindlists/api/mocks/auth"
	repoMocks "github.com/leggettc18/grindlists/api/mocks/pg"
	"github.com/leggettc18/grindlists/api/pg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateList(t *testing.T) {
	auth := new(authMocks.AuthService)
	repo := new(repoMocks.Repository)
	app, err := app.New()
	assert.Nil(t, err)
	c := client.New(handler.NewDefaultServer(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			Repository: repo,
			App:        *app,
			Auth:       auth,
		},
	})))
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
