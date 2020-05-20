package main

import (
	"file"
	"types"
	"os"
	"Response"
	"strings"
	"dataOp"
	"norm"
	"gonum.org/v1/gonum/mat"
	"math"
	"fmt"
)

const weigthSize = 10
const FileName = "datasets/houses.csv"
var sVector = 0
var sMatrix = 0

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
	dataOp.FormatData(&Train_Data, Data, 1)
	Pred := types.PredictD{}
	predict(Train_Data, &Pred)
}

func predict(TR types.DataTrain, P *types.PredictD) {

	var header []string
	var datasets []float64

	Sc := [4]string{"Gryffindor", "Hufflepuff", "Ravenclaw", "Slytherin"}
	P.Weights = FormateWeigths(file.ReadDat("datasets/weights.csv"))
	GetProb(P, Sc, TR)
	header = append(header, "Index", "Hogwarts House")
	file.SaveHeader(header, FileName)
	for i := 0; i < len(P.Probas[0]); i++ {
		datasets = append(datasets, float64(i), 0.0)
	}
	fmt.Println(datasets)
	file.SaveLines(datasets, FileName)
}

func GetProb(P *types.PredictD, Sc [4]string, TR types.DataTrain) {

	var res mat.VecDense

	P.Probas = make(map[int][]float64)
	sVector = len(TR.Line[0])
	sMatrix = len(TR.Line)
	trainMat := mat.NewDense(sMatrix, sVector, Tranform(TR.Line))

	for i := 0; i < len(Sc); i++ {
		theta := mat.NewVecDense(weigthSize + 1, P.Weights[i])
		res.MulVec(trainMat, theta)
		z := g(res)
		P.Probas[i] = z.RawVector().Data
	}
}

func g(z mat.VecDense) (mat.VecDense) {

	var final *mat.VecDense

	data := z.RawVector().Data
	final = mat.NewVecDense(sMatrix, gtab(data))
	return (*final)
}

func gtab(data []float64) ([]float64) {

	for i := 0; i < len(data); i++ {
		data[i] = 1 / (1 + math.Exp(-1 * data[i]))
	}
	return (data)
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

func FormateWeigths(data []float64) (map[int][]float64) {

	tab := make(map[int][]float64)

	for i := 0; i < len(data); i += weigthSize + 1 {
		tab[len(tab)] = data[i : i + weigthSize + 1]
	}
	return (tab)
}