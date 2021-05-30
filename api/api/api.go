package api

import "github.com/leggettc18/grindlists/api/app"

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