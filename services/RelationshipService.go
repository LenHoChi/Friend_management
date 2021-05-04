package services

import (
	// "context"
	// "fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	// "Friend_management/db"
	// "Friend_management/models"
	r_AddFriend "Friend_management/models/request"
	"Friend_management/repository"
	// "encoding/json"
	"fmt"
	// "github.com/gin-gonic/gin"
)

var RelationshipKey = "relationshipKey"

func Relationship (router chi.Router) {
	router.Get("/", getAllRelationships)
	router.Post("/", makeFriend)
	router.Post("/len", FindListFriend)
}

func getAllRelationships (w http.ResponseWriter, r *http.Request) {
	relationships, err := repository.GetAllRelationship(dbInstance)
	if err != nil {
		responseWithJSON(w, http.StatusOK ,"bug")
		return
	}
	if err := render.Render(w, r, relationships); err != nil {
		responseWithJSON(w, http.StatusOK, "bug")
		return
	}

}

func makeFriend(w http.ResponseWriter, r *http.Request){
	requestAddFriend := &r_AddFriend.AddRelationship{}
	render.Bind(r, requestAddFriend)
	userEmail := requestAddFriend.Friends[0]
	fmt.Println("emai ne: ",requestAddFriend.Friends[0])
	friendEmail := requestAddFriend.Friends[1]
	if err := repository.AddRelationship(dbInstance, userEmail, friendEmail); err != nil {
		return
	}
	responseWithJSON(w, http.StatusOK, "success")
}

func FindListFriend(w http.ResponseWriter, r *http.Request){
	requestFriendList := &r_AddFriend.RequestFriendListx{}
	render.Bind(r, requestFriendList)
	fmt.Println("emai ne: ",requestFriendList.Email)
	lst, _ := repository.FindListFriend(dbInstance, requestFriendList.Email)
	render.Render(w, r, lst)
}



