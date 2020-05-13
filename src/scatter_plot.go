package main

import (
	"os"
	"file"
	"types"
	"norm"
	"Response"
	"strings"
	"dataOp"
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
	res := file.ReadFile(&Data, args[0], 0)
	if res != 0 {
		return
	}
	norm.NormalizeAllData(&Data)
	Res := types.SaveCor{}
	dataOp.ScatterPlot(Data, &Res)
	fmt.Println(Res.Cor)
	fmt.Println(Res.Name1)
	fmt.Println(Res.Name2)
}