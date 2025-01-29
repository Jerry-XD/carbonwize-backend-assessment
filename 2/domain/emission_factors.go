package domain

import "time"

type EmissionFactors struct {
	ID             uint      `gorm:"primaryKey;column:id;autoIncrement:true"`
	Name           string    `gorm:"column:name;type:varchar(255)"`
	Detail         string    `gorm:"column:detail;type:varchar(1000)"`
	ActivityType   string    `gorm:"column:activity_type;type:varchar(255)"`
	UnitType       string    `gorm:"column:unit_type;type:varchar(255)"`
	VehicleType    string    `gorm:"column:vehicle_type;type:varchar(255)"`
	FuelType       string    `gorm:"column:fuel_type;type:varchar(255)"`
	EmissionFactor float64   `gorm:"column:emission_factor;type:float"`
	Source         string    `gorm:"column:source;type:varchar(255)"`
	CreatedAt      time.Time `gorm:"column:created_at;type:time"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:time"`
}

func (EmissionFactors) TableName() string {
	return "emission_factors"
}
