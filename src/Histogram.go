package main

import (
	"os"
	"file"
	"types"
	"norm"
	"Response"
	"strings"
	"dataOp"
	"graph"
	"fmt"
	"images"
	"strconv"
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
	if verbose == true {
		Response.PrintVerboseStep("[1/3] Normalize all data ...")
		Response.PrintVerbose("All data are converted between 0 and 1")
	}
	norm.NormalizeAllData(&Data)
	if verbose == true {
		Response.PrintVerboseStep("[2/3] Get all subject ...")
		Response.PrintVerbose("subject are : Gryffindor, Hufflepuff, Ravenclaw, Slytherin")
	}
	dataOp.GetMat(Data.Categ, Data.Mat)
	GraphV := dataOp.Calc(Data)
	Tab := [4]string{"Gryffindor", "Hufflepuff", "Ravenclaw", "Slytherin"}
	if verbose == true {
		Response.PrintVerboseStep("[3/3] Calcul value and draw graph ...")
		
		str := fmt.Sprintf("-------------------------------------\n")
		str += fmt.Sprintf("Calcul : we need Average and Standard deviation\n")
		str += fmt.Sprintf("Standard deviation graphical representation : \n")
		str += fmt.Sprintf("| 2.1 %s | 13.6 %s | 34.1 %s | 34.1 %s | 13.6 %s | 2.1 %s | \n", "%", "%", "%", "%", "%", "%")
		str += fmt.Sprintf("-------------------------------------\n")
		
		Response.PrintVerbose(str)
	}
	z := 0
	for i := 0; i < len(GraphV.Table); i += 4 {
		graph.DrawOne(GraphV, Tab, Data.Mat[z], i, verbose)
		z++
	}

	Response.Sucess("Append image ...")

	for a := 0; a < 13; a++ {

		if a % 2 == 0 && a != len(Data.Mat) - 1 {
			images.Append("graphs/test" + strconv.Itoa(a) + ".png", "graphs/" + Data.Mat[a] + ".png", "graphs/" + Data.Mat[a + 1] + ".png")
		}
	}
	
	images.Append("graphs/test3.png", "graphs/test0.png", "graphs/test2.png")
	images.Append("graphs/test1.png", "graphs/test4.png", "graphs/test6.png")
	images.Append("graphs/test2.png", "graphs/test8.png", "graphs/test10.png")

	images.AppendRow("graphs/test0.png", "graphs/test3.png", "graphs/test1.png")
	images.AppendRow("graphs/test3.png", "graphs/test2.png", "graphs/test0.png")
	images.AppendRow("graphs/histogram.png", "graphs/" + Data.Mat[len(Data.Mat) - 1] + ".png", "graphs/test3.png")

	Response.Sucess("Drawing histogram on term ...")
	images.DrawOnTerm("graphs/histogram.png")
}