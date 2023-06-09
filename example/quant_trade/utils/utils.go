package utils

import (
	"crypto/rand"
	"fmt"
)

func CalPercentLevels(gridNum int) []float64 {
	if gridNum <= 0 {
		panic(fmt.Sprintf("invalid gridNum(%d)", gridNum))
	}
	total := 1.0
	grid := total / float64(gridNum)
	var levels []float64
	for i := 0; i < gridNum; i++ {
		levels = append(levels, float64(i)*grid)
	}
	levels = append(levels, 1.0)
	return levels
}

func CalPriceLevels(highest float64, lowest float64, gridNum int) []float64 {
	if gridNum <= 0 {
		panic(fmt.Sprintf("invalid gridNum(%d)", gridNum))
	}
	if lowest < 0 || highest <= lowest {
		panic(fmt.Sprintf("invalid highest(%.2f) lowest(%.2f)", highest, lowest))
	}
	total := highest - lowest
	grid := total / float64(gridNum)
	var levels []float64
	for i := 0; i < gridNum; i++ {
		levels = append(levels, highest-float64(i)*grid)
	}
	levels = append(levels, lowest)
	return levels
}

func GetUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return BytesToHex(b)
}

func BytesToHex(b []byte) string {
	return fmt.Sprintf("%X", b)
}
