package responses

type CalculateResponse struct {
	ActivityType     string  `gorm:"column:activity_type" json:"activity_type"`
	CarbonEmissionKG float64 `gorm:"column:carbon_emission_kg" json:"carbon_emission_kg"`
}
