package models

import ("net/http")

type RequestFriendListx struct{
	Email string `json:"email"`
}
type RequestFriendLists struct {
	RequestFriendLists []RequestFriendListx `json:"friends"`
}
func (u *RequestFriendListx) Bind(r *http.Request) error{
	return nil
}
func (*RequestFriendLists) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// type AddRelationship2 struct {
// 	Friends []string `json:"friends"`
// }

// func (u *AddRelationship2) Bind(r *http.Request) error {
// 	return nil
// }