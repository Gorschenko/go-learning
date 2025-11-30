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

func (repository *LinksRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	// repository.Database.DB.First(&link, "hash = ? OR id = ?", hash, id)
	result := repository.Database.DB.First(&link, "hash = ?", hash)

	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}
