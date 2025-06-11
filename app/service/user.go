package service

import (
	"kajiLabTeam/DEBUKATU/lib"
	"kajiLabTeam/DEBUKATU/model"
	"log"
)

type UserService struct{}

func (UserService) GetUsers(userId int64) ([]model.User, error) {
	users := []model.User{}

	query := lib.DB.Model(&model.User{}).Where("deleted = ?", false)
	if userId > 0 {
		query = query.Where("user_id = ?", userId)
	}

	if err := query.Find(&users).Error; err != nil {
		log.Printf("DBクエリエラー: %v", err)
		return nil, err
	}

	return users, nil
}

func (UserService) CreateUser(name string) (int64, error) {
	log.Printf("lib.DB is nil? %v\n", lib.DB == nil)

	user := model.User{
		Name:    name,
		Deleted: false,
	}

	if err := lib.DB.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.UserId, nil
}

func (UserService) UpdateUser(userId int64, userName string) error {
	db := lib.DB

	return db.Model(&model.User{}).
		Where("user_id = ? AND deleted = ?", userId, false).
		Update("name", userName).Error
}
