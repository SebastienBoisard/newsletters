package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const DeleteTableBlock = `
DROP TABLE IF EXISTS Block;
`
const CreateTableBlock = `
CREATE TABLE Block (
  NewsletterID char(27)     NOT NULL,
  EpisodeID    int,
  BlockID      int          NOT NULL,
  Content      text         NOT NULL,
  PRIMARY KEY (NewsletterID, EpisodeID, BlockID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
`

type Block struct {
	NewsletterID string `db:"NewsletterID" json:"newsletterID"`
	EpisodeID    int    `db:"EpisodeID"    json:"episodeID"`
	BlockID      int    `db:"BlockID"      json:"blockID"`
	Content      string `db:"Content"      json:"content"`
}

var errBlocksNotFound = errors.New("blocks not found")

func GetBlocks(db *sqlx.DB, newsletterId string, episodeID int) (*[]Block, error) {

	log.Printf("GetBlocks - BEGIN")
	log.Printf("GetBlocks - newsletterId=%s", newsletterId)
	log.Printf("GetBlocks - episodeID=%d", episodeID)

	blocks := make([]Block, 0)
	err := db.Select(&blocks, "SELECT * FROM Block WHERE NewsletterId=? AND EpisodeID=?", newsletterId, episodeID)
	if err != nil {
		log.Printf("GetBlocks - Error while retrieving blocks [err=%v]", err)
		return nil, errBlocksNotFound
	}

	log.Printf("GetBlocks - blocks=%+v", blocks)
	log.Printf("GetBlocks - END")

	return &blocks, nil
}
