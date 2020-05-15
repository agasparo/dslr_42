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
)

func main() {


	//inputs -> moyenne des notes de chaques eleves de chaques classes
	//	moyene >= 0.5 ok
	//	moyenne < 0.5 pas ok
	//	Replace NaN data with median
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
	Learn := types.Learning{ 0.3, 1, 0.0, 0.0, make(map[int]float64) }
	fmt.Println(Learn)
	Train(Train_Data, &Learn)
	fmt.Println(Learn)
}

func Train(Tr types.DataTrain, Learn *types.Learning) {

	for i := 0; i < len(Tr.Data); i++ {

		theta := GradientDescent(Tr, Learn, i)
		Learn.Weights[len(Learn.Weights)] = theta
	}
}



func GradientDescent(Tr types.DataTrain, Learn *types.Learning, index int) (float64) {

	length := len(Tr.Data[index].Line)

	for i := 0; i < Learn.MaxIterations; i++ {

	}
	return (0.0)
}