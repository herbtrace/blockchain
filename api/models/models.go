package models

import "time"

// LatLong represents geographical coordinates
type LatLong struct {
	Lat     float64 `json:"lat"`
	Long    float64 `json:"long"`
	Address string  `json:"address,omitempty"`
}

// EnvironmentalConditions represents farming environment data
type EnvironmentalConditions struct {
	SoilQuality       string  `json:"soil_quality,omitempty"`
	Moisture          float64 `json:"moisture,omitempty"`
	Temperature       float64 `json:"temperature,omitempty"`
	Humidity          float64 `json:"humidity,omitempty"`
	WeatherConditions string  `json:"weather_conditions,omitempty"`
	IrrigationMethod  string  `json:"irrigation_method,omitempty"`
}

// FarmingInputs represents inputs used in farming
type FarmingInputs struct {
	Fertilizers      string `json:"fertilizers,omitempty"`
	PesticidesUsed   string `json:"pesticides_used,omitempty"`
	OrganicCertified bool   `json:"organic_certified"`
}

// PermitCompliance represents regulatory compliance
type PermitCompliance struct {
	PermitID   string `json:"permit_id"`
	PermitType string `json:"permit_type"`
	Issuer     string `json:"issuer"`
	ValidUntil string `json:"valid_until,omitempty"`
}

// CollectionEvent represents the collection event data
type CollectionEvent struct {
	BatchID     string                  `json:"batch_id" binding:"required"`
	ActorID     string                  `json:"actor_id" binding:"required"`
	CropID      string                  `json:"crop_id" binding:"required"`
	Location    LatLong                 `json:"location" binding:"required"`
	StartDate   time.Time               `json:"start_date" binding:"required"`
	HarvestDate time.Time               `json:"harvest_date" binding:"required"`
	Environment EnvironmentalConditions `json:"environment,omitempty"`
	Inputs      FarmingInputs           `json:"inputs,omitempty"`
	Permits     []PermitCompliance      `json:"permits,omitempty"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
