package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const DeleteTableHeader = `
DROP TABLE IF EXISTS Header;
`
const CreateTableHeader = `
CREATE TABLE Header (
  HeaderID     char(27)     NOT NULL,
  Content      text         NOT NULL,
  PRIMARY KEY (HeaderID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
`

type Header struct {
	HeaderID string `db:"HeaderID"     json:"headerID"`
	Content  string `db:"Content"      json:"content"`
}

var errHeaderNotFound = errors.New("header not found")

func GetHeader(db *sqlx.DB, headerID string) (*Header, error) {

	log.Printf("GetHeader - BEGIN")
	log.Printf("GetHeader - headerID=%s", headerID)

	header := Header{}
	err := db.Get(&header, "SELECT * FROM Header WHERE HeaderID=?", headerID)
	if err != nil {
		log.Printf("GetHeader - Error while retrieving header [err=%v]", err)
		return nil, errHeaderNotFound
	}

	log.Printf("GetHeader - header=%+v", header)
	log.Printf("GetHeader - END")

	return &header, nil
}
