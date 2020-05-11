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
}