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

type Learning struct {

	LearningRate 	float64
	MaxIterations	int
	Theta0			float64
	Theta1			float64
	LengthK			float64
	LengthP			float64
	Perte			float64
}