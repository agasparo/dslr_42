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

func Variance(data map[int]float64) (float64) {

	var n, Sum, SumSq float64

	for i := 0; i < len(data); i++ {

		n = float64(i + 1)
		Sum += data[i]
		SumSq += math.Pow(data[i], 2)
	}
	V := (SumSq - (Sum * Sum) / n) / (n - 1)
	return (V)
}

func Covariance(x, y map[int]float64) (float64) {

	var n, Sum, xm, ym, cx, cy float64

	cx = Count(x)
	cy = Count(y)

	xm = Mean(cx, x)
	ym = Mean(cy, y)

	for i := 0; i < len(x); i++ {

		n = float64(i + 1)
		Sum += (x[i] - xm) * (y[i] - ym)

	}
	return (Sum / n)
}

func Correlation(x, y map[int]float64) (float64) {

	var covar, etx, ety, cx, cy float64

	if len(x) != len(y) {
		return (-1.0)
	}

	covar = Covariance(x, y)
	cx = Count(x)
	cy = Count(y)
	etx = Std(Mean(cx, x), cx, x)
	ety = Std(Mean(cy, y), cy, y)

	return (covar / (etx * ety))
}

func Percent(per float64, data map[int]float64) (float64) {

	C := Count(data)
	return (C * per / 100)
}

func MaptoSlice(data map[int]float64) ([]float64) {

	var Tab []float64

	for i := 0; i < len(data); i++ {
		Tab = append(Tab, data[i])
	}
	return (Tab)
}

func SliceToMap(data []float64) (map[int]float64) {

	Tab := make(map[int]float64)

	for i := 0; i < len(data); i++ {
		Tab[len(Tab)] = data[i]
	}
	return (Tab)
}