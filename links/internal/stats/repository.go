package stats

import (
	"test/packages/db"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
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

func (repository *StatsRepository) GetStats(from time.Time, to time.Time, by string) []GetStatsRespose {
	var stats []GetStatsRespose
	var selectQuery string

	switch by {
	case "day":
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	case "month":
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks)"
	}
	query := repository.Database.DB.
		Table("stats").
		Select(selectQuery).
		Session(&gorm.Session{})

	if false {
		query.Where("sum > 10")
	}

	query.
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)
	return stats
}
