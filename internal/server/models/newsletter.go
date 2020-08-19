package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const DeleteTableNewsletter = `
DROP TABLE IF EXISTS Newsletter;
`

const CreateTableNewsletter = `
CREATE TABLE Newsletter (
  NewsletterID 	char(27) 	 NOT NULL,
  Name 		    varchar(64)  DEFAULT NULL,
  PRIMARY KEY (NewsletterID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
`

type Newsletter struct {
	NewsletterID string `db:"NewsletterID"       json:"newsletterId"`
	Name         string `db:"Name"     json:"name"`
}

var errNewsletterNotFound = errors.New("newsletter not found")

func GetNewsletter(db *sqlx.DB, newsletterId string) (*Newsletter, error) {

	log.Printf("getNewsletter - BEGIN")
	log.Printf("getNewsletter - newsletterId=%s", newsletterId)

	newsletter := Newsletter{}
	err := db.Get(&newsletter, "SELECT * FROM Newsletter WHERE NewsletterID=?", newsletterId)
	if err != nil {
		log.Printf("getNewsletter - Error while retrieving newsletter [err=%v]", err)
		return nil, errNewsletterNotFound
	}

	log.Printf("getNewsletter - newsletter=%+v", newsletter)
	log.Printf("getNewsletter - END")

	return &newsletter, nil
}
