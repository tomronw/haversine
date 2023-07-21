package haversine

import (
	"math"
)

func returnDistance(lat1, lon1, lat2, lon2 float32) float32 {
	lat1, lon1, lat2, lon2 = dtr(lat1), dtr(lon1), dtr(lat2), dtr(lon2)

	dlat := lat2 - lat1
	dlon := lon2 - lon1
	a := math.Sin(float64(dlat/2))*math.Sin(float64(dlat/2)) + math.Cos(float64(lat1))*math.Cos(float64(lat2))*math.Sin(float64(dlon/2))*math.Sin(float64(dlon/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	// 6371 is the radius of the Earth in km btw
	return float32(6371.0) * float32(c)
}
func dtr(degrees float32) float32 {
	return degrees * math.Pi / 180
}
func IsInRadius(lat1, lon1, lat2, lon2 float32, limit int) bool {
	return returnDistance(lat1, lon1, lat2, lon2) <= float32(limit)
}
