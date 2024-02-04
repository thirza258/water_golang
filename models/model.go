package models

import (
	"gorm.io/gorm"
	"time"
)

type Quality struct {
	gorm.Model
	City         string "json:city"
	Region       string "json:region"
	Country      string "json:country"
	AirQuality   int    "json:air_quality"
	WaterQuality int    "json:water_quality"
}

func (q *Quality) BeforeCreate(tx *gorm.DB) (err error) {
	q.CreatedAt = time.Now()
	return
}
