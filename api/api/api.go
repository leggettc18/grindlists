package api

import (
	"github.com/gorilla/mux"
	"github.com/leggettc18/grindlists/api/app"
	"github.com/leggettc18/grindlists/api/auth"
	"github.com/leggettc18/grindlists/api/gqlgen"
	"github.com/leggettc18/grindlists/api/pg"
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

func (api *API) Init(r *mux.Router) {
	db, err := pg.Open(api.App.Config.Server.DbConn)
	if err != nil {
		panic(err)
	}
	repo := pg.NewRepository(db)
	authSvc := auth.NewAuth(api.App.Config.SecretKey, api.App.Config.SecretKey)
	r.Handle("/", gqlgen.NewPlaygroundHandler("/graphql"))
	r.Handle("/graphql", gqlgen.NewHandler(repo, *api.App, authSvc))
	r.Use(authSvc.AuthMiddleware)
}
