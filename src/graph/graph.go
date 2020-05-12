package graph

import (
	"types"
	"github.com/wcharczuk/go-chart"
	"os"
	"fmt"
	"maths"
	"maps"
)

func DrawOne(Data types.GraphHisto, Sc [4]string, name string, mat string) {

	All := make(map[int]int)

	for i := 0; i < 4; i++ {
		min := maths.Min(Data.Stats)
		key := maps.ArraySearchFloat(Data.Stats, min)
		All[i] = key
		Data.Stats[key] = maths.Max(Data.Stats) + Data.Stats[key]
	}
	fmt.Println(All)
	DrawHisto(Data, Sc, All, name, mat)
}

func DrawHisto(Data types.GraphHisto, Sc [4]string, All map[int]int, name string, mat string) {
	
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
			{Value: maths.Count(Data.Table[All[0]].Cat), Label: Sc[All[0]]},
			{Value: maths.Count(Data.Table[All[1]].Cat), Label: Sc[All[1]]},
			{Value: maths.Count(Data.Table[All[2]].Cat), Label: Sc[All[2]]},
			{Value: maths.Count(Data.Table[All[3]].Cat), Label: Sc[All[3]]},
		},
	}


	f, _ := os.Create("graphs/" + name)
	defer f.Close()
	graph.Render(chart.PNG, f)
	//os.Remove("output.png")
}