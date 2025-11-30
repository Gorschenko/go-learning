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

func (repository *LinksRepository) Create(link *Link) (*Link, error) {
	result := repository.Database.DB.Create(link)

	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}
