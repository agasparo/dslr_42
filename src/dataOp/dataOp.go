package dataOp

import (
	"types"
	"math"
	"fmt"
	"maths"
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

func ScatterPlot(data types.Datas, Resp *types.SaveCor) {

	min := GetMin(data)
	min = 1215
	ReduceData(&data, int(min))
	for cols := 6; cols < len(data.Table); cols++ {

		for n_cols := 6; n_cols < len(data.Table); n_cols++ {

			if n_cols > cols {
				cor := maths.Correlation(data.Table[cols].Cat, data.Table[n_cols].Cat)
				fmt.Printf("Correlation between %s | %s : %f\n", data.Categ[cols], data.Categ[n_cols], cor)
				if math.Abs(cor) > Resp.Cor {
					Resp.Cor = math.Abs(cor)
					Resp.Name1 = data.Categ[cols]
					Resp.Name2 = data.Categ[n_cols]
					Resp.Map1 = data.Table[cols].Cat
					Resp.Map2 = data.Table[n_cols].Cat
				}
			}
		}
	}
}

func GetMin(data types.Datas) (float64) {

	min := maths.Count(data.Table[6].Cat)

	for i := 7; i < len(data.Table); i++ {
		c := maths.Count(data.Table[i].Cat)
		if c < min {
			min = c
		}
	}
	return (min)
}

func ReduceData(data *types.Datas, max int) {

	for i := 6; i < len(data.Table); i++ {

		N := types.Dat{}
		N.Cat = Reduce(data.Table[i].Cat, max)
		data.Table[i] = N
	}
}

func Reduce(data map[int]float64, max int) (map[int]float64) {

	tab := make(map[int]float64)

	for i := 0; i < max; i++ {
		tab[len(tab)] = data[i]
	}
	return (tab)
}