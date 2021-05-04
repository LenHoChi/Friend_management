package repository

import (
	// "database/sql"
	"Friend_management/db"
	"Friend_management/models"
	r_Request "Friend_management/models/request"

	// "container/list"
)

func GetAllRelationship(database db.Database) (*models.RelationshipList, error) {
	list := &models.RelationshipList{}

	rows, err := database.Conn.Query("SELECT * FROM relationship")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var relationship models.Relationship
		err := rows.Scan(&relationship.UserEmail, &relationship.FriendEmail, &relationship.AreFriend, &relationship.IsSubcriber, &relationship.IsBlock)
		if err != nil {
			return list, err
		}
		list.Relationships = append(list.Relationships, relationship)
	}
	return list, nil
}

func AddRelationship(database db.Database , userEmail string ,friendEmail string) error{
	query := `INSERT INTO relationship values ($1, $2, $3, $4, $5)`

	database.Conn.QueryRow(query, userEmail, friendEmail, true, false, false)
	return nil
}

func FindListFriend(database db.Database, email string) (*r_Request.RequestFriendLists, error){
	list := &r_Request.RequestFriendLists{}
	query := `select friend_email from relationship where user_email = $1 and arefriends = true
	 union
	 select user_email from relationship where friend_email = $1 and arefriends = true`

	rows, err := database.Conn.Query(query, email)

	if err != nil{
		return list, err
	}
	for rows.Next(){
		var requestFriendList r_Request.RequestFriendListx
		rows.Scan(&requestFriendList.Email)
		list.RequestFriendLists = append(list.RequestFriendLists, requestFriendList)
	}
	return list, nil
}