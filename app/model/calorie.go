package model

import "time"

type Calorie struct {
	CalorieId      int64     `gorm:"column:calorie_id;primaryKey;autoIncrement"`
	WeightId       int64     `gorm:"column:weight_id"`
	ModelId        int64     `gorm:"column:model_id"`
	MustCalorie    int64     `gorm:"column:must_calorie"`
	CurrentCalorie int64     `gorm:"column:current_calorie"`
	CreatedDate    time.Time `gorm:"column:created_date;type:datetime;autoCreateTime"`
}

func (Calorie) TableName() string { return "Calorie_Data" }
