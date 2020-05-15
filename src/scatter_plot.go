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
	"images"
)

func main() {

	var verbose bool
	
	args := os.Args[1:]
	if len(args) < 1 || len(args) > 2 {
		Response.Print("usage : ./describe [file_name.csv] [-v : verbose]")
		return
	}
	if len(args) == 2 && args[1] != "-v" {
		Response.Print("usage : ./describe [file_name.csv] [-v : verbose]")
		return
	}
	if len(args) == 2 && args[1] == "-v" {
		verbose = true
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
	dataOp.ScatterPlot(Data, &Res, verbose)
	Response.Sucess(fmt.Sprintf("Most : %s and %s with cor : %f\n", Res.Name1, Res.Name2, Res.Cor))
	graph.DrawScatterPlot(Res)
	images.DrawOnTerm("graphs/ScatterPlot.png")
}