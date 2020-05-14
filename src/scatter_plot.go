package main

import (
	"os"
	"file"
	"types"
	"Response"
	"strings"
	"dataOp"
	"fmt"
	"graph"
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
	Res := types.SaveCor{}
	dataOp.ScatterPlot(Data, &Res)
	fmt.Printf("Most : %s and %s with cor : %f\n", Res.Name1, Res.Name2, Res.Cor)
	graph.DrawScatterPlot(Res)
}