package dataOp

import (
	"types"
	"maths"
	"math"
)

func GetMat(cat map[int]string, mat map[int]string) {

	for i := 6; i < len(cat); i++ {
		mat[len(mat)] = cat[i]
	}
}

func Calc(data types.Datas) (types.GraphHisto) {

	Save := types.GraphHisto{}
	Save.Table = make(map[int]types.Dat)
	Sc := [4]string{"Gryffindor", "Hufflepuff", "Ravenclaw", "Slytherin"}

	for i := 0; i < len(data.Mat); i++ {

		Dat := types.Dat{}
		Dat.Cat = make(map[int]float64)
		for z := 0; z < 4; z++ {
			Dat.Cat = Getdatas(Sc[z], i + 6, data.Table, data.School)
			Save.Table[len(Save.Table)] = Dat
		}
	}
	return (Save)
}

func Getdatas(search string, index int, data map[int]types.Dat, school map[int]string) (map[int]float64) {

	tmp := make(map[int]float64)

	for z := 0; z < len(data[0].Cat); z++ {
		if school[z] == search && !math.IsNaN(data[index].Cat[z]) {
			tmp[len(tmp)] = data[index].Cat[z]
		}
	}
	return (tmp)
}

func CountSearch(data map[int]float64, percent int, add int, ref_min float64) (float64, float64) {

	var per, count float64
	var total int

	count = maths.Count(data)
	per = maths.Percentile(float64(add) * count, percent, data)
	total = CountPerc(ref_min, per, data)
	return float64(total), per
}

func CountPerc(min float64, max float64, data map[int]float64) (int) {

	var c int

	for i := 0; i < len(data); i++ {

		if data[i] <= max && data[i] >= min && data[i] != min {
			c++
		}
		if data[i] == 0 {
			c++
		}
	}
	return (c)
}
