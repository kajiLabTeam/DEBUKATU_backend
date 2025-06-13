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
