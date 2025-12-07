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
	today := datatypes.Date(time.Now())
	repository.Database.DB.Find(&stat, "link_id = ? and date = ?", linkId, today)

	if stat.ID == 0 {
		repository.Database.DB.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   today,
		})
	} else {
		stat.Clicks++
		repository.Database.DB.Save(&stat)
	}
}
