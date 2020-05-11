package maths

import (
	"math"
	"sort"
)

func Max(data map[int]float64) (float64) {

	max := data[0]

	for i := 0; i < len(data); i++ {
		
		if data[i] > max {
			max = data[i]
		}
	}
	return (max)
}

func Min(data map[int]float64) (float64) {

	min := data[0]

	for i := 0; i < len(data); i++ {

		if data[i] < min {
			min = data[i]
		}
	}
	return (min)
}

func Count(data map[int]float64) (float64) {

	c := 0

	for i := 0; i < len(data); i++ {
		c++
	}
	return (float64(c))
}

func Mean(c float64, data map[int]float64) (float64) {

	var res float64

	for i := 0; i < len(data); i++ {
		res += data[i]
	}
	return (res / c)
}

func Std(m float64, c float64, data map[int]float64) (float64) {
	
	var sd float64

	for i := 0; i < len(data); i++{
		
		sd += math.Pow((data[i] - m), 2)
	}
	sd = math.Sqrt(sd / c)
	return (sd)
}

func Percentile(c float64, divi int, data map[int]float64) (float64) {

	SortTable := MaptoSlice(data)
	sort.Float64s(SortTable)
	index := int(math.Ceil(c / float64(divi)) - 1)
	for in, element := range SortTable {
		if in == index {
			return (element)
		}
	}
	return (0.0)
}

func MaptoSlice(data map[int]float64) ([]float64) {

	var Tab []float64

	for i := 0; i < len(data); i++ {
		Tab = append(Tab, data[i])
	}
	return (Tab)
}