package dataOp

import (
	"types"
	//"fmt"
)

func GetMat(cat map[int]string, mat map[int]string) {

	for i := 6; i < len(cat); i++ {
		mat[len(mat)] = cat[i]
	}
}

func Calc(data types.Datas) {

	Save := types.GraphHisto{}
	Save.Table = make(map[int]types.Dat)

	for i := 0; i < len(data.Mat); i++ {

		Dat := types.Dat{}
		Dat.Cat = make(map[int]float64)
		for z := 0; z < len(data.School); z++{
			Getdatas(data.School[z], i + 6)
		}
		Save.Table[len(Save.Table)] = Dat
		return
	}
}

func Getdatas(search string, index int) {

}