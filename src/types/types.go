package types

import (

)

type Datas struct {

	Table 	map[int]Dat
	Categ	map[int]string
	School	map[int]string
	Mat 	map[int]string
}

type Dat struct {

	Cat map[int]float64
}

type GraphHisto struct {

	Table map[int]Dat
	Stats map[int]float64
}

type SaveCor struct {

	Name1 string
	Name2 string
	Map1 	map[int]float64
	Map2	map[int]float64
	Cor 	float64
}

type DataTrain struct {

	Moy	map[int]float64
}