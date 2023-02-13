package entity

type Coordinate struct {
	Latitude  float64 `gorm:"comment:纬度"`
	Longitude float64 `gorm:"comment:经度"`
}
