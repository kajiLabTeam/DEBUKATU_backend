package service

import (
	"kajiLabTeam/DEBUKATU/lib"
	"kajiLabTeam/DEBUKATU/model"
	"log"
)

type WeightService struct{}

func (WeightService) GetWeights(userId int64, modelId int64) ([]model.Weight, error) {
	weights := []model.Weight{}

	if err := lib.DB.Model(&model.Weight{}).Where("user_id = ?", userId).Where("model_id=?", modelId).Find(&weights).Error; err != nil {
		log.Printf("DBクエリエラー: %v", err)
		return nil, err
	}

	return weights, nil
}

func (WeightService) CreateWeight(userId int64, modelId int64, weight float64) (int64, error) {
	log.Printf("lib.DB is nil? %v\n", lib.DB == nil)

	model := model.Weight{
		ModelId:       modelId,
		UserId:        userId,
		CurrentWeight: weight,
	}

	if err := lib.DB.Create(&model).Error; err != nil {
		return 0, err
	}

	return model.WeightId, nil
}
