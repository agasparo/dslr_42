package main

import (
	"file"
	"types"
	"os"
	"Response"
	"strings"
	"dataOp"
	"norm"
	"fmt"
)

const weigthSize = 10

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

	//proba
	//weigth
	weigth := FormateWeigths(file.ReadDat("datasets/weights.csv"))
	fmt.Println(weigth)
	//Sc := [4]string{"Gryffindor", "Hufflepuff", "Ravenclaw", "Slytherin"}
	//fmt.Println(Sc)
}

func FormateWeigths(data []float64) (map[int][]float64) {

	tab := make(map[int][]float64)

	for i := 0; i < len(data); i += weigthSize + 1 {
		tab[i] = data[i : i + weigthSize + 1]
		fmt.Println(tab[i])
		fmt.Println("------------------------------------")
	}
	return (tab)
}