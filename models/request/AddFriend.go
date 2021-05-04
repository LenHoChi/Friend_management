package models

import "net/http"
 type AddRelationship struct {
	 Friends []string `json:"friends"`
 }

 func (u *AddRelationship) Bind(r *http.Request) error {
	 return nil
 }