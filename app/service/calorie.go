package service

import (
	"kajiLabTeam/DEBUKATU/lib"
	"kajiLabTeam/DEBUKATU/model"
	"log"
)

type CalorieService struct{}

func (CalorieService) GetCalories(weightId int64) ([]model.Calorie, error) {
	calories := []model.Calorie{}

	if err := lib.DB.Model(&model.Calorie{}).Where("weight_id = ?", weightId).Find(&calories).Error; err != nil {
		log.Printf("DBクエリエラー: %v", err)
		return nil, err
	}

	return calories, nil
}

func (CalorieService) CreateCalorie(weightId int64, modelId int64, mustCalorie int64, currentCalorie int64) (int64, error) {
	log.Printf("lib.DB is nil? %v\n", lib.DB == nil)

	calorie := model.Calorie{
		WeightId:       weightId,
		ModelId:        modelId,
		MustCalorie:    mustCalorie,
		CurrentCalorie: currentCalorie,
	}

	if err := lib.DB.Create(&calorie).Error; err != nil {
		return 0, err
	}

	return calorie.CalorieId, nil
}
