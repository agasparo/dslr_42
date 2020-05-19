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

/*
	y tab de 0 et de 1
		pour nous 0 ou 1 cest ecole ou pas ecole

	cost function

		si y = 1 cost : -log(h(x))
		si y = 0 cost : -log(1 - h(x))

		value : Cost(h(x), y) = y  * log(h(x)) + (1 - y) * log(1 - h(x))
		h(x) = g(theta * x)
		g(x) = 1 / (1 - e(-x))

	gradientDescent

		z = matrice | data * theta car h(x) = g(theta * x)
		cacul du cout de la fonction
		si le cout - l'ancien cout < stop
			on stop l'ago
		on calcul le nouveau theta

			gradient 
				(z - y) * data / length
			calcul
				gradient * learning rate

 */

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
	theta := mat.NewVecDense(10, Trans(0.0, 10, 1))
	trainMat := mat.NewDense(len(Train_Data.Line), len(Train_Data.Line[0]), Tranform(Train_Data.Line))
	Learn := types.Learning{ 0.1, 1, theta, 1.0, 0.000001, make(map[int][]float64), trainMat } // pour iteration apres 100000
	Train(Train_Data, &Learn, Data)
}

func Train(Tr types.DataTrain, Learn *types.Learning, Data types.Datas) {

	var y map[int]float64

	Table := [4]string{"Gryffindor", "Hufflepuff", "Ravenclaw", "Slytherin"} 

	for i := 0; i < len(Table); i++ {

		y = dataOp.RempY(Table[i], Data.School)
		gradientDescent(Tr, Learn, y)
		Learn.Weights[len(Learn.Weights)] = Learn.Theta.RawVector().Data
		return
	}
}

func gradientDescent(Tr types.DataTrain, Learn *types.Learning, y map[int]float64) {

	var length, ac_cost float64
	var vec, z, sub mat.VecDense
	var mul, divi mat.Dense

	length = float64(len(y))
	length_mat := mat.NewDense(1, 10, Trans(length, 1, 10))

	for i := 0; i < Learn.MaxIterations; i++ {

		fmt.Println("--------------------------------------------------------")
		fmt.Printf("turn : %d\n", i)

		vec.MulVec(Learn.Datas, Learn.Theta)
		z = g(vec)
		ac_cost = Learn.Cost
		Learn.Cost = Cost(z, length, y)
		
		fmt.Printf("ac_cost : %f\n", ac_cost)
		fmt.Printf("Learn.Cost : %f\n", Learn.Cost)
		fmt.Printf("Learn.Stop : %f\n", Learn.Stop)
		if ac_cost - Learn.Cost < Learn.Stop {
			fmt.Println("la")
			break
		}

		sub.SubVec(&z, MaptoVec(y))
		mul.Mul(VecToMat(sub, 1, 1600), Learn.Datas)
		divi.DivElem(&mul, length_mat)
		Learn.Theta = mat.NewVecDense(10, GetGrad(divi.RawMatrix().Data, Learn.LearningRate, Learn.Theta.RawVector().Data))
		fmt.Println(Learn.Theta)
	}
}

func GetGrad(data []float64, rate float64, sub []float64) ([]float64) {

	for i := 0; i < len(data); i++ {
		data[i] = sub[i] - (rate * data[i])
	}
	return (data)
}

func VecToMat(vec mat.VecDense, x, y int) (*mat.Dense) {

	return (mat.NewDense(x, y, vec.RawVector().Data))
}

func Cost(z mat.VecDense, length float64, y map[int]float64) (float64) {

	var Sum, zLog float64

	zLog = LogVector(z)
	i := 0
	//for i := 0; i < len(y) ; i++ {

		Sum += y[i] * math.Log(zLog) + (1 - y[i]) * math.Log(1 - zLog)
	//}
	//Sum = -1 * (Sum / length)
	return (Sum)
}

func LogVector(vec mat.VecDense) (float64) {

	var Sum float64

	data := vec.RawVector().Data
	for i := 0; i < len(data); i++ {
		Sum += data[i]
	}
	return (Sum)
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

func MaptoVec(y map[int]float64) (*mat.VecDense) {

	var vector *mat.VecDense
	var tmp []float64

	for i := 0; i < len(y); i++ {
		tmp = append(tmp, y[i])
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