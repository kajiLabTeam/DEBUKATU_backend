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
func (UserService) GetUser(name string, password string) ([]model.User, error) {
	users := []model.User{}

	query := lib.DB.Model(&model.User{}).Where("deleted = ?", false)

	query = query.Where("name = ?", name).Where("password=?", password)

	if err := query.Find(&users).Error; err != nil {
		log.Printf("DBクエリエラー: %v", err)
		return nil, err
	}

	return users, nil
}

func (UserService) CreateUser(name string, password string, age int64, height float64) (int64, error) {
	log.Printf("lib.DB is nil? %v\n", lib.DB == nil)

	user := model.User{
		Name:     name,
		Password: password,
		Age:      age,
		Height:   height,
		Deleted:  false,
	}

	if err := lib.DB.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.UserId, nil
}

func (UserService) UpdateUser(userId int64, name string, deleted *bool) error {
	db := lib.DB

	updates := map[string]interface{}{}
	if name != "" {
		updates["name"] = name
	}
	if deleted != nil {
		updates["deleted"] = *deleted
	}

	if len(updates) == 0 {
		return nil // 何も更新しない
	}

	return db.Model(&model.User{}).
		Where("user_id = ? AND deleted = false", userId).
		Updates(updates).Error
}
