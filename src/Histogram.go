package main

import (
	"os"
	"file"
	"types"
	//"norm"
	"Response"
	"strings"
	"dataOp"
	"graph"
	"strconv"
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
	//norm.NormalizeAllData(&Data)
	dataOp.GetMat(Data.Categ, Data.Mat)
	GraphV := dataOp.Calc(Data)
	Tab := [4]string{"Ravenclaw", "Slytherin", "Gryffindor", "Hufflepuff"}
	z := 0
	i := 0
	//for i := 0; i < len(GraphV.Table); i += 4 {
		graph.DrawOne(GraphV, Tab, "outpout" + strconv.Itoa(z) + ".png", Data.Mat[z], i)
		z++
	//}
}