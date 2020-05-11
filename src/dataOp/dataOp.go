package dataOp

import (
	"types"
	//"fmt"
	//"maths"
)

func GetMat(cat map[int]string, mat map[int]string) {

	for i := 6; i < len(cat); i++ {
		mat[len(mat)] = cat[i]
	}
}

func Calc(data types.Datas) (types.GraphHisto) {

	Save := types.GraphHisto{}
	Save.Table = make(map[int]types.Dat)
	Sc := [4]string{"Ravenclaw", "Slytherin", "Gryffindor", "Transfiguration"}

	for i := 0; i < len(data.Mat); i++ {

		Dat := types.Dat{}
		Dat.Cat = make(map[int]float64)
		for z := 0; z < 4; z++ {
			Dat.Cat = Getdatas(Sc[z], i + 6, data.Table, data.School)
			Save.Table[len(Save.Table)] = Dat
		}
		return(Save)
	}
	return (Save)
}

func Getdatas(search string, index int, data map[int]types.Dat, school map[int]string) (map[int]float64) {

	tmp := make(map[int]float64)

	for z := 0; z < len(data[0].Cat); z++ {
		if school[z] == search {
			tmp[len(tmp)] = data[index].Cat[z]
		}
	}
	return (tmp)
}

/*func GetValFor(Data types.GraphHisto, begin int) {

	var res, tmp map[int]float64

	for i := begin; i < begin + 4; i++ {

		res = tmp
		for z := 0; z < len(Data.Table[i].Cat); z++ {
			Data.Table[i].Cat[z] = (Data.Table[i].Cat[z] - maths.Mean(Data.Table[i].Cat[z])) / maths.Std(Data.Table[i].Cat[z])	
		}
	}
	fmt.Println(Data.Table[0])
	fmt.Println(Data.Table[1])
	fmt.Println(Data.Table[2])
	fmt.Println(Data.Table[3])
	return
}*/