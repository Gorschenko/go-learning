package stats

import (
	"test/packages/db"
	"time"

	"gorm.io/datatypes"
)

type StatsRepository struct {
	Database *db.Db
}

func NewStatsRepository(database *db.Db) *StatsRepository {
	return &StatsRepository{
		Database: database,
	}
}

func (repository *StatsRepository) AddClick(linkId uint) {
	var stat Stat
	repository.Database.DB.Find(&stat, "link_id = ? and date = ?", datatypes.Date(time.Now()))

	if stat.ID == 0 {
		repository.Database.DB.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   datatypes.Date(time.Now()),
		})
	} else {
		stat.Clicks++
		repository.Database.DB.Save(&stat)
	}
}
