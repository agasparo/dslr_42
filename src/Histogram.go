package main

import (
	"github.com/wcharczuk/go-chart"
	"os"
	"file"
	"types"
	"norm"
	"Response"
	"strings"
	"fmt"
	"dataOp"
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
	norm.NormalizeAllData(&Data)
	dataOp.GetMat(Data.Categ, Data.Mat)
	fmt.Println(Data.School)
	fmt.Println(Data.Mat)
}

func Darw() {
	graph := chart.BarChart{
		Title: "Histogram",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		Bars: []chart.Value{ // cours 
			{Value: 5.25, Label: "Blue"},
			{Value: 4.88, Label: "Green"},
			{Value: 4.74, Label: "Gray"},
			{Value: 3.22, Label: "Orange"},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
	//os.Remove("output.png")
}