package main

import (
	"os"
	"file"
	"types"
	"Response"
	"strings"
	"dataOp"
	"graph"
	"images"
	"say"
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
	name, err := dataOp.PairPlot(Data)
	if err == 1 {
		return
	}
	Response.Sucess("Calculing pair plot ...")
	graph.DrawPairPlot(name)
	os.Remove(name)
	Response.Sucess("Drawing histogram on term ...")
	images.DrawOnTerm("graphs/pair_plot.png")
	say.PairPlot()
}