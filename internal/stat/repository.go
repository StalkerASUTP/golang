package stat

import (
	"go/adv-api/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.Db.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)
	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			Clicks: 1,
			LinkId: linkId,
			Date:   currentDate,
		})

	} else {
		stat.Clicks += 1
		repo.Db.Save(&stat)
	}
}

func (repo *StatRepository) GetStat(by string, from, to time.Time) []GetStatResponce {
	var stats []GetStatResponce
	var selectQuery string
	switch by {
	case GroupByDay:
		selectQuery = "to_char(date,'YYYY-MM-DD') as period, sum(clicks)"
	case GroupByMonth:
		selectQuery = "to_char(date,'YYYY-MM') as period, sum(clicks)"
	}
	repo.DB.Table("stats").
		Select(selectQuery).
		Where("date between ? and ?", from, to).
		Group("period").
		Order("period").Scan(&stats)
	return stats
}
