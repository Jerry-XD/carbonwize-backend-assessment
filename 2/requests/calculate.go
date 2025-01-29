package requests

type CalculateRequest struct {
	ActivityType string  `json:"activity_type"`
	DistanceKM   float64 `json:"distance_km"`
	VehicleType  string  `json:"vehicle_type"`
	FuelType     string  `json:"fuel_type"`
}
