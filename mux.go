package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"mikke-server/config"
	"mikke-server/database"
	"mikke-server/handler"
	"mikke-server/tools/clock"
	"mikke-server/usecase"
	"net/http"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	r := database.Repository{Clocker: clock.RealClocker{}}
	v := validator.New()
	db, cleanup, err := database.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	ps := &handler.PostQuestion{
		Usecase:   usecase.PostUsecase{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/post", ps.ServeHTTP)
	return mux, cleanup, err
}