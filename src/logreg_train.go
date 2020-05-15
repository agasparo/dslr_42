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
	//"math"
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
	Learn := types.Learning{ 0.3, 1, 0.0, 1.0, 0.000001, make(map[int]float64) } // pour iteration apres 100000
	Train(Train_Data, &Learn, Data)
	fmt.Println(Learn)
}

func Train(Tr types.DataTrain, Learn *types.Learning, Data types.Datas) {

	var y map[int]float64

	Table := [4]string{"Gryffindor", "Hufflepuff", "Ravenclaw", "Slytherin"} 

	for i := 0; i < len(Table); i++ {

		y = dataOp.RempY(Table[i], Data.School)
		gradientDescent(Tr, Learn, y)
		Learn.Weights[len(Learn.Weights)] = Learn.Theta
		return
	}
}

func gradientDescent(Tr types.DataTrain, Learn *types.Learning, y map[int]float64) {

	var length, ac_cost float64
	var tmpMat, z mat.Dense

	length = float64(len(y))

	for i := 0; i < Learn.MaxIterations; i++ {

		trainMat := mat.NewDense(len(Tr.Line), len(Tr.Line[0]), Tranform(Tr.Line))   //Tr.Data * Learn.Theta
		thetaMat := mat.NewDense(len(Tr.Line[0]), len(Tr.Line), Trans(Learn.Theta, len(Tr.Line[0]), len(Tr.Line)))

		fc := mat.Formatted(trainMat, mat.Prefix("    "), mat.Squeeze())
		fmt.Printf("tr :%v\n", fc)
		fc1 := mat.Formatted(thetaMat, mat.Prefix("    "), mat.Squeeze())
		fmt.Printf("th :%v\n", fc1)
		tmpMat.Mul(trainMat, thetaMat)
		fc2 := mat.Formatted(&tmpMat, mat.Prefix("    "), mat.Squeeze())
		fmt.Printf("fi :%v\n", fc2)
		z = g(tmpMat)
		fmt.Println(z) //ici
		ac_cost = Learn.Cost
		//Learn.Cost = Cost(z, length, y)
		//if ac_cost - Learn.Cost < Learn.Stop {
		//	break
		//}
		//gradient = (z - y) * Tr.Data / length
		//Learn.Theta = gradient * Learn.LearningRate
	}
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

/*func Cost(z map[int]float64, length float64, y map[int]float64) (float64) {

	var Sum float64

	for i := 0; i < len(y) ; i++ {

		Sum += y[i] * math.Log(z) + (1 - y[i]) * math.Log(1 - z)
	}
	Sum = -1 * (Sum / length)
	return (Sum)
}
*/
func g(z mat.Dense) (mat.Dense) {

	return (1 / (1 + math.Exp(-z)))
}