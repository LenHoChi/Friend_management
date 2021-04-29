package models

import (

)

type Relationship struct {
	userEmail string `json"user_email"`
	friendEmail string `json"friend_email"`
	areFriend bool `json"arefriends"`
	isSubcriber bool `json"issubcriber"`
	isBlock bool `json"isblock"`
} 