package links

import (
	"test/packages/db"
)

type LinksRepository struct {
	Database *db.Db
}

func NewLinksRepository(database *db.Db) *LinksRepository {
	return &LinksRepository{
		Database: database,
	}
}

func (repository *LinksRepository) Create(link *Link) {

}
