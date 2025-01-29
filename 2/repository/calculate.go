package repository

import (
	"context"

	"carbonwize-backend-assessment/domain"
	"carbonwize-backend-assessment/requests"
	"carbonwize-backend-assessment/responses"

	"gorm.io/gorm"
)

type CalculateRepo interface {
	GetEmissionFactors(ctx context.Context, inp *requests.CalculateRequest) (result *responses.CalculateResponse, err error)
}

type Repository struct {
	DB *gorm.DB
}

func NewCalculateRepo(gormDb *gorm.DB) *Repository {
	return &Repository{
		DB: gormDb,
	}
}

func (repo *Repository) GetEmissionFactors(ctx context.Context, inp *requests.CalculateRequest) (result *responses.CalculateResponse, err error) {
	result = &responses.CalculateResponse{}

	err = repo.DB.WithContext(ctx).Table(domain.EmissionFactors{}.TableName()).
		Select("activity_type, CAST(emission_factor * ? AS FLOAT) AS carbon_emission_kg", inp.DistanceKM).
		Where("activity_type = ? AND vehicle_type = ? AND fuel_type = ?", inp.ActivityType, inp.VehicleType, inp.FuelType).
		Scan(&result).Error

	return result, err
}
