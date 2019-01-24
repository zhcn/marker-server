package geo

import (
	"fmt"
	"github.com/mmcloughlin/geohash"
)

const LESS_1_KM_PRECISION uint = 30
const LESS_3_KM_PRECISION uint = 25
const LESS_20_KM_PRECISION uint = 20
const LESS_20_KM_MASK int64 = ^((1 << 10) - 1)

func EncodeTest() {
	var i uint
	for i = 1; i < 13; i++ {
		fmt.Println(geohash.EncodeWithPrecision(10, 10, i))
		fmt.Println(geohash.EncodeIntWithPrecision(10, 10, i))
	}
}

func EncodeInt(lat, lng float64) uint64 {
	//30 bit precision : lat 15bit lon 15bit km error: +-0.61km
	return geohash.EncodeIntWithPrecision(lat, lng, LESS_1_KM_PRECISION)
}

// default use LESS_20_KM_MASK
func GetIndexBound(geoCode uint64) (lowerBound uint64, upperBound uint64) {
	lowerBound = uint64(int64(geoCode) & LESS_20_KM_MASK)
	upperBound = uint64(int64(geoCode)&LESS_20_KM_MASK) + (1 << 10)
	return
}
