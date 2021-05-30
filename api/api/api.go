package api

import (
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/leggettc18/grindlists/api/app"
)

type API struct {
	App *app.App
}

func New(a *app.App) (api *API, err error) {
	api = &API{App: a}
	if err != nil {
		return nil, err
	}
	return api, nil
}

type query struct {}

func (_ *query) Hello() string {
	return "Hello World!"
}

func (api *API) Init(r *mux.Router) {
	s := `
		type Query {
			hello: String!
		}
	`

	schema := graphql.MustParseSchema(s, &query{})
	r.Handle("/graphql", &relay.Handler{Schema: schema})
}