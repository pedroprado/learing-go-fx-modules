package serverfx

import (
	"context"
	"net/http"

	"example.auto.wiring/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(newMux),
	fx.Invoke(
		InitServer,
	),
)

func newMux(httpHandlers *handlers.HttpHandlers) *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/person/{id}", httpHandlers.PersonHandler.Get)
	mux.HandleFunc("/person", httpHandlers.PersonHandler.Create)
	return mux
}

func InitServer(lifecycle fx.Lifecycle, mux *mux.Router) {
	server := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Close()
		},
	})
}
