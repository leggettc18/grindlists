package gqlgen

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/leggettc18/grindlists/api/app"
	"github.com/leggettc18/grindlists/api/pg"
)

// NewHandler returns a new graphql endpoint handler.
func NewHandler(repo pg.Repository, app app.App) http.Handler {
	return handler.NewDefaultServer(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			Repository: repo,
			App: app,
		},
	}))
}

// NewPlaygroundHandler returns a new GraphQL Playground handler.
func NewPlaygroundHandler(endpoint string) http.Handler {
	return playground.Handler("GraphQL Playground", endpoint)
}