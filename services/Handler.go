package services

import (
	"net/http"

	"github.com/go-chi/chi"
	// "github.com/go-chi/render"
	"Friend_management/db"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db

	router.Route("/users", Users)
	router.Route("/relationship", Relationship)
	return router
}