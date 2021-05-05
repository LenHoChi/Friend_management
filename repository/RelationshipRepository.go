package repository

import (
	"Friend_management/db"
	"Friend_management/models"
	r_Response "Friend_management/models/response"
	"errors"
	"fmt"

	// "fmt"
	"strings"
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

func AddRelationship(database db.Database , userEmail string ,friendEmail string) (*r_Response.ResponseSuccess ,error){
	//check email similar
	//check relationship similar
	query := `INSERT INTO relationship values ($1, $2, $3, $4, $5)`
	// database.Conn.QueryRow(query, userEmail, friendEmail, true, false, false)
	_, err := database.Conn.Exec(query,userEmail, friendEmail, true, false, false)
	if err != nil{
		return nil, errors.New("Error: "+err.Error())
	}
	return &r_Response.ResponseSuccess{Success: true}, nil
}

func FindListFriend(database db.Database, email string) (*r_Response.ResponseListFriend, error){
	list := &r_Response.ResponseListFriend{}
	query := `select friend_email from relationship where user_email = $1 and arefriends = true
	 union
	 select user_email from relationship where friend_email = $1 and arefriends = true`

	rows, err := database.Conn.Query(query, email)

	if err != nil{
		return list, err
	}
	for rows.Next(){
		var email string
		err := rows.Scan(&email)
		if err != nil{
			return nil,nil
		}
		list.Friends = append(list.Friends, email)
	}
	list.Success = true
	list.Count = len(list.Friends)
	return list, nil
}

func FindCommonListFriend(database db.Database, lstEmail []string) (*r_Response.ResponseListFriend ,error) {
	list := &r_Response.ResponseListFriend{}
	//check same email
	query := `select r.friend_email from relationship r
	where r.user_email in ($1,$2) and r.arefriends =true 
	group by r.friend_email 
	having count(r.friend_email)>1`
	rows, err := database.Conn.Query(query, lstEmail[0], lstEmail[1])

	if err != nil {
		return list, nil
	}
	for rows.Next(){
		var email string
		err := rows.Scan(&email)
		if err != nil {
			return nil, nil
		}
		list.Friends = append(list.Friends, email)
	}
	list.Success = true
	list.Count = len(list.Friends)
	return list, nil
}

func BeSubcribe(database db.Database, requestor string, target string) (*r_Response.ResponseSuccess ,error) {
	query := `update relationship set issubcriber =true where user_email =$1 and friend_email =$2`
	// database.Conn.QueryRow(query, requestor, target)
	//check email same
	_, err := database.Conn.Exec(query, requestor, target)
	if err != nil{
		return nil, nil
	}
	return &r_Response.ResponseSuccess{Success: true}, nil
}

func ToBlock(database db.Database, requestor string, target string) (*r_Response.ResponseSuccess, error) {
	query := `update relationship set issubcriber =false where user_email=$1 and friend_email=$2`
	// database.Conn.QueryRow(query,requestor,target)
	//check email same
	_, err := database.Conn.Exec(query,requestor,target)
	if err != nil {
		return nil, nil
	}
	return &r_Response.ResponseSuccess{Success: true}, nil
}

func RetrieveUpdate(database db.Database, sender string, target string) (*r_Response.ResponseRetrieve, error){
	list := &r_Response.ResponseRetrieve{}
	query := `select friend_email from relationship 
	where user_email =$1 and (arefriends=true or issubcriber=true)
	and isblock =false`
	rows, err := database.Conn.Query(query, sender)
	if err != nil {
		return list, nil
	}
	for rows.Next(){
		var email string
		err := rows.Scan(&email)
		if err != nil{
			return nil, nil
		}
		list.Recipients = append(list.Recipients, email)
	}
	list.Success = true
	return list, nil
}
func SplitString(text string) string{
	split := strings.Split("ho chi len"," ")
	fmt.Sprintln(split[0])
	return "nil"
}