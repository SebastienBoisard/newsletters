package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
)

const DeleteTableEpisode = `
DROP TABLE IF EXISTS Episode;
`
const CreateTableEpisode = `
CREATE TABLE Episode (
  NewsletterID 	    char(27) 	 NOT NULL,
  EpisodeID         int,	
  CreationDate 		datetime     NOT NULL,
  HeaderID 	        char(27) 	 NOT NULL,
  FooterID 	        char(27) 	 NOT NULL,
  PRIMARY KEY (NewsletterID, EpisodeID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
`

type Episode struct {
	NewsletterID string    `db:"NewsletterID"    json:"newsletterId"`
	EpisodeID    int       `db:"EpisodeID"      json:"episodeID"`
	CreationDate time.Time `db:"CreationDate"    json:"creationDate"`
	HeaderID     string    `db:"HeaderID"    json:"headerID"`
	FooterID     string    `db:"FooterID"    json:"footerID"`
}

var errEpisodeNotFound = errors.New("episode not found")

func GetEpisode(db *sqlx.DB, newsletterID string, episodeID int) (*Episode, error) {

	log.Printf("GetEpisode - BEGIN")
	log.Printf("GetEpisode - newsletterID=%d", newsletterID)
	log.Printf("GetEpisode - episodeID=%d", episodeID)

	episode := Episode{}
	err := db.Get(&episode, "SELECT * FROM Episode WHERE NewsletterID = ? AND EpisodeID=?", newsletterID, episodeID)
	if err != nil {
		log.Printf("GetEpisode - Error while retrieving episode [err=%v]", err)
		return nil, errEpisodeNotFound
	}

	log.Printf("GetEpisode - episode=%+v", episode)
	log.Printf("GetEpisode - END")

	return &episode, nil
}
