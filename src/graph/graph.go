package graph

import (
	"types"
	"github.com/wcharczuk/go-chart"
	"os"
	"fmt"
	"maths"
	"dataOp"
	"maps"
)

func DrawOne(Data types.GraphHisto, Sc [4]string, name string, mat string, i int) {

	All := make(map[int]int)

	for i := 0; i < 4; i++ {
		min := maths.Min(Data.Stats)
		key := maps.ArraySearchFloat(Data.Stats, min)
		All[i] = key
		Data.Stats[key] = maths.Max(Data.Stats) + Data.Stats[key]
	}
	fmt.Println(All)
	DrawHisto(Data, Sc, All, name, mat, i)
}

func DrawHisto(Data types.GraphHisto, Sc [4]string, All map[int]int, name string, mat string, i int) {
	

	v1 := dataOp.CountSearch(Data.Table[i + All[0]].Cat, 4, 1, 0.0)
	v2 := dataOp.CountSearch(Data.Table[i + All[0]].Cat, 2, 1, v1)
	v3 := dataOp.CountSearch(Data.Table[i + All[0]].Cat, 4, 3, v2)
	v4 := dataOp.CountSearch(Data.Table[i + All[0]].Cat, 1, 1, v3)

	v5 := dataOp.CountSearch(Data.Table[i + All[1]].Cat, 4, 1, 0.0)
	v6 := dataOp.CountSearch(Data.Table[i + All[1]].Cat, 2, 1, v5)
	v7 := dataOp.CountSearch(Data.Table[i + All[1]].Cat, 4, 3, v6)
	v8 := dataOp.CountSearch(Data.Table[i + All[1]].Cat, 1, 1, v7)

	v9 := dataOp.CountSearch(Data.Table[i + All[2]].Cat, 4, 1, 0.0)
	v10 := dataOp.CountSearch(Data.Table[i + All[2]].Cat, 2, 1, v9)
	v11 := dataOp.CountSearch(Data.Table[i + All[2]].Cat, 4, 3, v10)
	v12 := dataOp.CountSearch(Data.Table[i + All[2]].Cat, 1, 1, v11)

	v13 := dataOp.CountSearch(Data.Table[i + All[3]].Cat, 4, 1, 0.0)
	v14 := dataOp.CountSearch(Data.Table[i + All[3]].Cat, 2, 1, v13)
	v15 := dataOp.CountSearch(Data.Table[i + All[3]].Cat, 4, 3, v14)
	v16 := dataOp.CountSearch(Data.Table[i + All[3]].Cat, 1, 1, v15)

	graph := chart.BarChart{
		Title: "Histogram " + mat,
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   256,
		BarWidth: 30,
		Bars: []chart.Value{
			{Value: v1, Label: Sc[All[0]]},
			{Value: v2, Label: Sc[All[0]]},
			{Value: v3, Label: Sc[All[0]]},
			{Value: v4, Label: Sc[All[0]]},
			
			{Value: v5, Label: Sc[All[1]]},
			{Value: v6, Label: Sc[All[1]]},
			{Value: v7, Label: Sc[All[1]]},
			{Value: v8, Label: Sc[All[1]]},

			{Value: v9, Label: Sc[All[2]]},
			{Value: v10, Label: Sc[All[2]]},
			{Value: v11, Label: Sc[All[2]]},
			{Value: v12, Label: Sc[All[2]]},

			{Value: v13, Label: Sc[All[3]]},
			{Value: v14, Label: Sc[All[3]]},
			{Value: v15, Label: Sc[All[3]]},
			{Value: v16, Label: Sc[All[3]]},
		},
	}


	f, _ := os.Create("graphs/" + name)
	defer f.Close()
	graph.Render(chart.PNG, f)
	//os.Remove("output.png")
}