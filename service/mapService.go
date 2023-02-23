package service

import "math"

// 返回值的单位为米
func EarthDistance(lat1, long1, lat2, long2 float64) float64 {
	radius := float64(6378137) // 6378137
	rad := math.Pi / 180.0

	lat1 = lat1 * rad
	long1 = long1 * rad
	lat2 = lat2 * rad
	long2 = long2 * rad
	theta := long2 - long1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}
