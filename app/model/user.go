package model

type User struct {
	UserId   int64   `gorm:"column:user_id;primaryKey;autoIncrement"`
	Name     string  `gorm:"column:name"`
	Password string  `gorm:"column:password"`
	Age      int64   `gorm:"column:age"`
	Height   float64 `gorm:"column:height"`
	Gender   string  `gorm:"column:gender"`
	Deleted  bool    `gorm:"column:deleted"`
}

func (User) TableName() string { return "User_Data" }
