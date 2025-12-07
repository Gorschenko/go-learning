package links

import (
	"test/packages/db"

	"gorm.io/gorm/clause"
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

func (repository *LinksRepository) Update(link *Link) (*Link, error) {
	result := repository.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repository *LinksRepository) Delete(id uint) error {
	result := repository.Database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *LinksRepository) GetById(id uint) (*Link, error) {
	var link Link
	result := repository.Database.DB.First(&link, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repository *LinksRepository) Count(limit, offset int) (int64, error) {
	var count int64

	result := repository.Database.DB.
		Table("links").
		Where("deleted_at IS NULL").
		Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (repository *LinksRepository) GetAll(limit, offset int) ([]*Link, error) {
	var links []*Link

	result := repository.Database.DB.
		Table("links").
		Where("deleted_at IS NULL").
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&links)

	if result.Error != nil {
		return nil, result.Error
	}

	return links, nil
}
