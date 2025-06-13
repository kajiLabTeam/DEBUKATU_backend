package model

import "time"

type Weight struct {
	WeightId      int64     `gorm:"column:weight_id;primaryKey;autoIncrement"`
	ModelId       int64     `gorm:"column:model_id"`
	UserId        int64     `gorm:"column:user_id"`
	CurrentWeight float64   `gorm:"column:model_weight"`
	CreatedDate   time.Time `gorm:"column:created_date;type:datetime;autoCreateTime"`
}

func (Weight) TableName() string { return "Weight_Data" }
