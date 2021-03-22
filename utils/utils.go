package utils

import (
	"log"
	"strconv"
)

func StringToFloat(value string) float64 {
	fValue := value
	var s float64
	s, err := strconv.ParseFloat(fValue, 64)
	if err != nil {
		log.Println(err)
	}

	return s
}

func StringToUint(value string) uint64 {
	fValue := value
	var s uint64
	s, err := strconv.ParseUint(fValue, 10, 64)
	if err != nil {
		log.Println(err)
	}

	return s
}
