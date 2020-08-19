package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const DeleteTableFooter = `
DROP TABLE IF EXISTS Footer;
`
const CreateTableFooter = `
CREATE TABLE Footer (
  FooterID     char(27)     NOT NULL,
  Content      text         NOT NULL,
  PRIMARY KEY (FooterID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
`

type Footer struct {
	FooterID string `db:"FooterID"     json:"footerID"`
	Content  string `db:"Content"      json:"content"`
}

var errFooterNotFound = errors.New("footer not found")

func GetFooter(db *sqlx.DB, footerID string) (*Footer, error) {

	log.Printf("GetFooter - BEGIN")
	log.Printf("GetFooter - footerID=%s", footerID)

	footer := Footer{}
	err := db.Get(&footer, "SELECT * FROM Footer WHERE FooterID=?", footerID)
	if err != nil {
		log.Printf("GetFooter - Error while retrieving footer [err=%v]", err)
		return nil, errFooterNotFound
	}

	log.Printf("GetFooter - footer=%+v", footer)
	log.Printf("GetFooter - END")

	return &footer, nil
}
