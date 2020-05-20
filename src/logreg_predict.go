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
	predict(Train_Data)
}

func predict(TR types.DataTrain) {

	Sc := [4]string{"Gryffindor", "Hufflepuff", "Ravenclaw", "Slytherin"} 
	fmt.Println(TR)
}