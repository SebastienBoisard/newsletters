package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const DeleteTableSubscription = `DROP TABLE IF EXISTS Subscription;`

const CreateTableSubscription = `
CREATE TABLE Subscription (
  NewsletterID  char(27)     NOT NULL,
  SubscriberID  char(27)     NOT NULL,
  NewsletterShortname varchar(32) DEFAULT NULL,
  NewsletterKey char(5) NOT NULL,
  SubscriberKey char(5) NOT NULL,
  StartEpisodeID 	int,
  EndEpisodeID   	int,
  PRIMARY KEY (NewsletterID, SubscriberID, NewsletterShortname)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
`

type Subscription struct {
	NewsletterID        string `db:"NewsletterID" json:"newsletterID"`
	SubscriberID        string `db:"SubscriberID" json:"subscriberID"`
	NewsletterShortname string `db:"NewsletterShortname" json:"newsletterShortname"`
	NewsletterKey       string `db:"NewsletterKey" json:"newsletterKey"`
	SubscriberKey       string `db:"SubscriberKey" json:"subscriberKey"`
	StartEpisodeID      int    `db:"StartEpisodeID" json:"startEpisodeID"`
	EndEpisodeID        int    `db:"EndEpisodeID" json:"endEpisodeID"`
}

var errSubscriptionNotFound = errors.New("subscription not found")

func GetSubscription(db *sqlx.DB, newsletterShortname string, episodeID int, newsletterKey string, subscriberKey string) (*Subscription, error) {

	log.Printf("GetSubscription - BEGIN")
	log.Printf("GetSubscription - newsletterShortname=%s", newsletterShortname)
	log.Printf("GetSubscription - episodeID=%d", episodeID)
	log.Printf("GetSubscription - newsletterKey=%s", newsletterKey)
	log.Printf("GetSubscription - subscriberKey=%s", subscriberKey)

	subscription := Subscription{}
	err := db.Get(&subscription, "SELECT * FROM Subscription WHERE NewsletterShortname=? AND NewsletterKey=? AND SubscriberKey=?",
		newsletterShortname, newsletterKey, subscriberKey)
	if err != nil {
		log.Printf("GetSubscription - Error while retrieving subscription [err=%v]", err)
		return nil, errSubscriptionNotFound
	}

	log.Printf("GetSubscription - subscription=%+v", subscription)
	log.Printf("GetSubscription - END")

	return &subscription, nil
}
