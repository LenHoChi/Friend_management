package repository

import (
	"database/sql"
	"Friend_management/models"
	// "fmt"
	"Friend_management/db"
)

func GetAllUsers(database db.Database) (*models.UserList, error) {
	list := &models.UserList{}

	rows, err := database.Conn.Query("SELECT * FROM users")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Email)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}
func AddUser(database db.Database, user *models.User) error {
	query := `INSERT INTO users (email) VALUES ($1)`
	// err := database.Conn.QueryRow(query, user.Email)
	// if err != nil {
	// 	return fmt.Errorf("have problem while insert")
	// }
	// return nil

	database.Conn.QueryRow(query, user.Email)
	return nil
}

func GetUserByEmail(database db.Database, email string) (models.User, error){
	user := models.User{}

	query := `select * from users where email = $1;`

	err := database.Conn.QueryRow(query, email).Scan(&user.Email)
	if err != nil{
		if err == sql.ErrNoRows{
			return user, err
		}
		return user, err
	}
	return user, nil
}

func DeleteUser(database db.Database, email string) error {
	query := `delete from users where email =$1`
	_, err := database.Conn.Exec(query, email)
	switch err {
		case sql.ErrNoRows:
			return db.ErrNoMatch
		default:
			return err
	}
}