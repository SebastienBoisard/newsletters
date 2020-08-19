package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const DeleteTableUser = `
DROP TABLE IF EXISTS User;
`

const CreateTableUser = `
CREATE TABLE User (
  UserID 		char(27) 	 NOT NULL,
  LastName 		varchar(64)  DEFAULT NULL,
  FirstName 	varchar(64)  DEFAULT NULL,
  Email 		varchar(128) DEFAULT NULL,
  PRIMARY KEY (UserID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
`

type User struct {
	UserID    string `db:"UserID"       json:"userId"`
	LastName  string `db:"LastName"     json:"lastName"`
	FirstName string `db:"FirstName"    json:"firstName"`
	Email     string `db:"Email"        json:"email"`
}

var errUserNotFound = errors.New("user not found")

func GetUser(db *sqlx.DB, userId string) (*User, error) {

	log.Printf("GetUser - BEGIN")
	log.Printf("GetUser - userId=%s", userId)

	user := User{}
	err := db.Get(&user, "SELECT * FROM User WHERE UserID=?", userId)
	if err != nil {
		log.Printf("GetUser - Error while retrieving campaign [err=%v]", err)
		return nil, errUserNotFound
	}

	log.Printf("GetUser - user=%+v", user)
	log.Printf("GetUser - END")

	return &user, nil
}
