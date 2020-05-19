package main

import (
	"file"
	"types"
	"os"
	"Response"
	"strings"
	"dataOp"
	"fmt"
	"norm"
	"math"
	"gonum.org/v1/gonum/mat"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		Response.Print("usage : ./describe [file_name.csv]")
		return
	}
	ext := strings.Split(args[0], ".")
	if ext[len(ext) -1] != "csv" {
		Response.Print("Your file must be a .csv")
		return
	}
	Data := types.Datas{}
	res := file.ReadFile(&Data, args[0], 1)
	if res != 0 {
		return
	}
	Data = dataOp.SupprUs([]int{0, 1, 2, 3, 4, 5, 6, 9, 16}, Data)
	norm.NormalizeAllData(&Data)
	Train_Data := types.DataTrain{}
	dataOp.FormatData(&Train_Data, Data)
	theta := mat.NewVecDense(11, Trans(0.0, 11, 1))
	trainMat := mat.NewDense(len(Train_Data.Line), len(Train_Data.Line[0]), Tranform(Train_Data.Line))
	Learn := types.Learning{ 0.1, 100000, theta, 1.0, 0.000001, make(map[int][]float64), trainMat }
	Train(Train_Data, &Learn, Data)
}

func Train(Tr types.DataTrain, Learn *types.Learning, Data types.Datas) {

	var y map[int]float64

	Table := [4]string{"Gryffindor", "Hufflepuff", "Ravenclaw", "Slytherin"} 

	for i := 0; i < len(Table); i++ {

		y = dataOp.RempY(Table[i], Data.School)
		gradientDescent(Tr, Learn, y)
		Learn.Weights[len(Learn.Weights)] = Learn.Theta.RawVector().Data
	}
	file.SaveFile(Learn.Weights, "datasets/weights.csv")
}

func gradientDescent(Tr types.DataTrain, Learn *types.Learning, y map[int]float64) {

	var length, ac_cost float64
	var vec, z, sub mat.VecDense
	var mul, divi mat.Dense

	length = float64(len(y))
	length_mat := mat.NewDense(1, 11, Trans(length, 1, 11))

	for i := 0; i < Learn.MaxIterations; i++ {

		vec.MulVec(Learn.Datas, Learn.Theta)
		z = g(vec)
		ac_cost = Learn.Cost
		Learn.Cost = Cost(z, length, y)
		
		if ac_cost - Learn.Cost < Learn.Stop {
			break
		}

		sub.SubVec(&z, MaptoVec(y, 0))
		mul.Mul(VecToMat(sub, 1, 1600), Learn.Datas)
		divi.DivElem(&mul, length_mat)
		grad := mat.NewVecDense(11, GetGrad(divi.RawMatrix().Data, Learn.LearningRate))
		Learn.Theta.SubVec(Learn.Theta, grad)
	}
}

func GetGrad(data []float64, rate float64) ([]float64) {

	for i := 0; i < len(data); i++ {
		data[i] = rate * data[i]
	}
	return (data)
}

func VecToMat(vec mat.VecDense, x, y int) (*mat.Dense) {

	return (mat.NewDense(x, y, vec.RawVector().Data))
}

func Cost(z mat.VecDense, length float64, y map[int]float64) (float64) {

	var Sum float64
	var res, res1, res2 mat.VecDense

	z0 := mat.NewVecDense(1600, LogVector(z, 0))
	z1 := mat.NewVecDense(1600, LogVector(z, 1))
	y1 := MaptoVec(y, 0)
	y2 := MaptoVec(y, 1)

	res.MulElemVec(y1, z0)
	res1.MulElemVec(y2, z1)
	res2.AddVec(&res, &res1)
	Sum = Summ(res2)
	Sum = Sum / -length
	return (Sum)
}

func Summ(vec mat.VecDense) (S float64) {

	data := vec.RawVector().Data
	for i := 0; i < len(data); i++ {
		S += data[i]
	}
	return (S)
}

func LogVector(vec mat.VecDense, add int) ([]float64) {

	var tab []float64

	data := vec.RawVector().Data
	for i := 0; i < len(data); i++ {
		if add == 0 {
			tab = append(tab, math.Log(data[i]))
		} else {
			tab = append(tab, math.Log(1 - data[i]))
		}
	}
	return (tab)
}

func g(z mat.VecDense) (mat.VecDense) {

	var final *mat.VecDense

	data := z.RawVector().Data
	final = mat.NewVecDense(1600, gtab(data))
	return (*final)
}

func gtab(data []float64) ([]float64) {

	for i := 0; i < len(data); i++ {
		data[i] = 1 / (1 + math.Exp(-1 * data[i]))
	}
	return (data)
}

func MaptoVec(y map[int]float64, add int) (*mat.VecDense) {

	var vector *mat.VecDense
	var tmp []float64

	for i := 0; i < len(y); i++ {
		if add == 0 {
			tmp = append(tmp, y[i])
		} else {
			tmp = append(tmp, 1.0 - y[i])
		}
	}

	vector = mat.NewVecDense(1600, tmp)
	return (vector)
}

func Trans(z float64, sizec, sizel int) (res []float64) {

	for i := 0; i < sizec; i++ {

		for a := 0; a < sizel; a++ {
			res = append(res, z)
		}
	}
	return (res)
}

func Tranform(data map[int]map[int]float64) ([]float64) {

	var Tab []float64

	for i := 0; i < len(data); i++ {

		for a := 0; a < len(data[i]); a++ {
			Tab = append(Tab, data[i][a])
		}
	}
	return (Tab)
}