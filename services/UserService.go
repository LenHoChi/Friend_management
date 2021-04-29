package services

import (
	"context"
	"fmt"
	"net/http"
	// "strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"Friend_management/db"
	"Friend_management/models"
	"Friend_management/repository"

	"encoding/json"
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
		responseWithJSON(w, http.StatusOK, "bug for ?")
		return
	}
	if err := render.Render(w, r, users); err != nil {
		responseWithJSON(w, http.StatusOK, "bug for ?")
		return 
	}
}

func createUser(w http.ResponseWriter, r *http.Request){
	user := &models.User{}

	if err := render.Bind(r, user); err != nil {
		responseWithJSON(w, http.StatusOK, "Bug for 1")
		return
	}
	if err:= repository.AddUser(dbInstance, user); err != nil {
		responseWithJSON(w, http.StatusOK, err.Error())
		return 
	}
	if err := render.Render(w, r, user); err != nil {
		responseWithJSON(w, http.StatusOK, "Bug for 3")
		return     
	}
}

func UserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		email := chi.URLParam(r, "emailID")
		if email == "" {
			responseWithJSON(w, http.StatusOK, "Email is required")
			return
		}
		fmt.Println("day: ",email)
		// id, err := strconv.Atoi(email)
		// if err != nil {
		// 	responseWithJSON(w, http.StatusOK, "invalid user mail")
		// }
		// fmt.Println("day2: ",id)
		ctx := context.WithValue(r.Context(), UserEmailKey, email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("vao toi day roi ne")
	email := r.Context().Value(UserEmailKey).(string)
	fmt.Println("neeee: ",email)
	user, err := repository.GetUserByEmail(dbInstance, email)
	if err != nil {
		if err == db.ErrNoMatch {
			responseWithJSON(w, http.StatusOK, "No any user match")
		}else {
			responseWithJSON(w, http.StatusOK, err.Error())
		}
		return
	}
	if err := render.Render(w, r, &user); err != nil {
		responseWithJSON(w, http.StatusOK, "Bug for 2")
		return 
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value(UserEmailKey).(string)
	err := repository.DeleteUser(dbInstance, email)
	if err != nil {
		if err == db.ErrNoMatch {
			responseWithJSON(w, http.StatusOK, "No any user match")
		}else{
			responseWithJSON(w, http.StatusOK, err.Error())
		}
		return
	}
}
func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}){
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}