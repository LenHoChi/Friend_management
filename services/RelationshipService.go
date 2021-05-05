package services

import (
	// "context"
	// "fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	// "Friend_management/db"
	// "Friend_management/models"
	r_Request "Friend_management/models/request"
	"Friend_management/repository"
	// "encoding/json"
	// "github.com/gin-gonic/gin"
)

var RelationshipKey = "relationshipKey"

func Relationship (router chi.Router) {
	router.Get("/", getAllRelationships)
	router.Post("/make", makeFriend)
	router.Post("/list", FindListFriend)
	router.Post("/common", FindCommonListFriend)
	router.Post("/update",BeSubcriber)
	router.Post("/block", ToBLock)
	router.Post("/retrieve", RetrieveUpdate)
	router.Post("/a", a)
}
func a (w http.ResponseWriter, r *http.Request) {
	repository.SplitString("a")
}
func getAllRelationships (w http.ResponseWriter, r *http.Request) {
	relationships, err := repository.GetAllRelationship(dbInstance)
	if err != nil {
		responseWithJSON(w, http.StatusOK ,"bug1")
		return
	}
	if err := render.Render(w, r, relationships); err != nil {
		responseWithJSON(w, http.StatusOK, "bug2")
		return
	}

}
// {"friends":["1","2"]}
func makeFriend(w http.ResponseWriter, r *http.Request){
	requestAddFriend := &r_Request.RequestFriendLists{}
	render.Bind(r, requestAddFriend)
	//check length <2
	userEmail := requestAddFriend.RequestFriendLists[0]
	friendEmail := requestAddFriend.RequestFriendLists[1]
	//check valid for two emails
	responseRS, err := repository.AddRelationship(dbInstance, userEmail, friendEmail)
	if err != nil {
		responseWithJSON(w, http.StatusBadRequest ,err.Error())
	}
	render.Render(w, r, responseRS)
}
//{"email":"1"}
func FindListFriend(w http.ResponseWriter, r *http.Request){
	Argument := &r_Request.RequestEmail{}
	render.Bind(r, Argument)
	responseRS, _ := repository.FindListFriend(dbInstance, Argument.Email)
	render.Render(w, r, responseRS)
}
// {"friends":["1","2"]}
func FindCommonListFriend(w http.ResponseWriter, r *http.Request){
	rsFriend := &r_Request.RequestFriendLists{}
	ls := make([]string,0)
	render.Bind(r, rsFriend)
	ls = append(ls, rsFriend.RequestFriendLists[0], rsFriend.RequestFriendLists[1])
	lst, _ := repository.FindCommonListFriend(dbInstance, ls)
	render.Render(w,r,lst)
}
// {"requestor":"len1","target":"len2"}
func BeSubcriber(w http.ResponseWriter, r *http.Request){
	Argument := &r_Request.RequestUpdate{}
	render.Bind(r, Argument)
	responseRS, _ := repository.BeSubcribe(dbInstance, Argument.Requestor, Argument.Target)
	render.Render(w, r, responseRS)
}
func ToBLock(w http.ResponseWriter, r *http.Request){
	Argument := &r_Request.RequestUpdate{}
	render.Bind(r, Argument)
	responseRS ,_ :=repository.ToBlock(dbInstance, Argument.Requestor, Argument.Target)
	render.Render(w, r, responseRS)
}
// {"sender":"len1","target":"len2"}
func RetrieveUpdate(w http.ResponseWriter, r *http.Request){
	Argument := &r_Request.RetrieveUpdate{}
	render.Bind(r, Argument)
	responseRS, _ := repository.RetrieveUpdate(dbInstance,Argument.Sender, Argument.Tartget)
	render.Render(w, r, responseRS)
}


