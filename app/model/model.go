package model

import "time"

type Model struct {
	ModelId       int64     `gorm:"column:model_id;primaryKey;autoIncrement"`
	UserId        int64     `gorm:"column:user_id"`
	ModelWeight   float64   `gorm:"column:model_weight"`
	LengthOfMonth float64   `gorm:"column:length_of_month"`
	CreatedDate   time.Time `gorm:"column:created_date;type:datetime;autoCreateTime"`
}

func (Model) TableName() string { return "Model_Data" }
