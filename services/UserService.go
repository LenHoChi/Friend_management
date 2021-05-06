package services

import (
	"context"
	"fmt"
	"net/http"
	// "strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	r_Response "Friend_management/models/response"

	"Friend_management/db"
	"Friend_management/models"
	"Friend_management/repository"

)

var UserEmailKey = "emailKey"

func Users(router chi.Router) {
	router.Get("/", getAllUsers)
	router.Post("/", createUser)

	router.Route("/{emailID}", func(router chi.Router) {
		router.Use(UserContext)
		router.Get("/", getUser)
		router.Delete("/", deleteUser)
	})
}
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.GetAllUsers(dbInstance)
	if err != nil {
		r_Response.ResponseWithJSON(w, http.StatusInternalServerError,err.Error())
		return
	}
	if err := render.Render(w, r, users); err != nil {
		r_Response.ResponseWithJSON(w, http.StatusInternalServerError,err.Error())
		return 
	}
}

func createUser(w http.ResponseWriter, r *http.Request){
	user := &models.User{}

	if err := render.Bind(r, user); err != nil {
		r_Response.ResponseWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err:= repository.AddUser(dbInstance, user); err != nil {
		r_Response.ResponseWithJSON(w, http.StatusInternalServerError, err.Error())
		return 
	}
	if err := render.Render(w, r, user); err != nil {
		r_Response.ResponseWithJSON(w, http.StatusInternalServerError, err.Error())
		return     
	}
}

func UserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		email := chi.URLParam(r, "emailID")
		if email == "" {
			r_Response.ResponseWithJSON(w, http.StatusInternalServerError, "Email is required")
			return
		}
		fmt.Println("day: ",email)
		ctx := context.WithValue(r.Context(), UserEmailKey, email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value(UserEmailKey).(string)
	user, err := repository.GetUserByEmail(dbInstance, email)
	if err != nil {
		if err == db.ErrNoMatch {
			r_Response.ResponseWithJSON(w, http.StatusInternalServerError, "No any user match")
		}else {
			r_Response.ResponseWithJSON(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	if err := render.Render(w, r, &user); err != nil {
		r_Response.ResponseWithJSON(w, http.StatusInternalServerError, err.Error())
		return 
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value(UserEmailKey).(string)
	err := repository.DeleteUser(dbInstance, email)
	if err != nil {
		if err == db.ErrNoMatch {
			r_Response.ResponseWithJSON(w, http.StatusInternalServerError, "No any user match")
		}else{
			r_Response.ResponseWithJSON(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
}

