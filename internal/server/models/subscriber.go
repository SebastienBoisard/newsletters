package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const DeleteTableSubscriber = `DROP TABLE IF EXISTS Subscriber;`

const CreateTableSubscriber = `
CREATE TABLE Subscriber (
  SubscriberID 	char(27) 	 NOT NULL,
  Email 		varchar(128) DEFAULT NULL,
  PRIMARY KEY (SubscriberID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
`

type Subscriber struct {
	SubscriberID string `db:"SubscriberID" json:"subscriberId"`
	Email        string `db:"Email"        json:"email"`
}

var errSubscriberNotFound = errors.New("subscriber not found")

func GetSubscriber(db *sqlx.DB, subscriberId string) (*Subscriber, error) {

	log.Printf("GetSubscriber - BEGIN")
	log.Printf("GetSubscriber - subscriberId=%s", subscriberId)

	subscriber := Subscriber{}
	err := db.Get(&subscriber, "SELECT * FROM Subscriber WHERE SubscriberID=?", subscriberId)
	if err != nil {
		log.Printf("GetSubscriber - Error while retrieving Subscriber [err=%v]", err)
		return nil, errSubscriberNotFound
	}

	log.Printf("GetSubscriber - subscriber=%+v", subscriber)
	log.Printf("GetSubscriber - END")

	return &subscriber, nil
}
