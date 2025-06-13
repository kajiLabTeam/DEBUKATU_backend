package service

import (
	"kajiLabTeam/DEBUKATU/lib"
	"kajiLabTeam/DEBUKATU/model"
	"log"
)

type ModelService struct{}

func (ModelService) GetModels(userId int64) ([]model.Model, error) {
	models := []model.Model{}

	if err := lib.DB.Model(&model.Model{}).Where("user_id = ?", userId).Find(&models).Error; err != nil {
		log.Printf("DBクエリエラー: %v", err)
		return nil, err
	}

	return models, nil
}

func (ModelService) CreateModel(userId int64, weight float64, month float64) (int64, error) {
	log.Printf("lib.DB is nil? %v\n", lib.DB == nil)

	model := model.Model{
		UserId:        userId,
		ModelWeight:   weight,
		LengthOfMonth: month,
	}

	if err := lib.DB.Create(&model).Error; err != nil {
		return 0, err
	}

	return model.ModelId, nil
}
