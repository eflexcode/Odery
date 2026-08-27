package database

import (
	"database/sql"

	"github.com/cmd/entity"
	_ "github.com/lib/pq"
)

func ConnectDatabase(url string) (*sql.DB, error) {

	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetUser(username string, db *sql.DB) (*entity.User, error) {

	var user entity.User

	rows, err := db.Query("SELECT * from users WHERE username = $1", username)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rows.Next()

	err = rows.Scan(user)
	
	if err != nil{
		return nil,err
	}
	
	return &user,nil
}
