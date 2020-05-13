package graph

import (
	"types"
	"maths"
	"sort"
	"os"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func DrawOne(Data types.GraphHisto, Sc [4]string, name string, mat string, g int) {

	var M, ET, C float64
	var Axe_X, Axe_Y []float64
	All := []chart.Series{}
	St := chart.Style{}

	for i := g % 4; i < 4; i++ {

		res := maths.MaptoSlice(Data.Table[g + i].Cat)
		sort.Float64s(res)
		nData := maths.SliceToMap(res)
		C = maths.Count(nData)
		M = maths.Mean(C, nData)
		ET = maths.Std(M, C, nData)

		Axe_X = append(Axe_X, 0.0, M - (2 * ET), M - ET, M, M + ET, M + (2 * ET), 1.0)
		Axe_Y = append(Axe_Y, 0.0, maths.Percent(2.1, nData), maths.Percent(13.6, nData), maths.Percent(34.1, nData), maths.Percent(13.6, nData), maths.Percent(2.1, nData), 1.0)

		if i == 0 {
			St = chart.Style {
					StrokeColor: drawing.ColorRed,              
					FillColor:   drawing.ColorRed.WithAlpha(64),
			}
		}

		if i == 1 {
			St = chart.Style{
					StrokeColor: drawing.ColorBlue,              
					FillColor:   drawing.ColorBlue.WithAlpha(64),
			}
		}

		if i == 2 {
			St = chart.Style{
					StrokeColor: drawing.ColorGreen,              
					FillColor:   drawing.ColorGreen.WithAlpha(64),
			}
		}

		if i == 3 {
			ColorCyan := drawing.Color{R: 0, G: 217, B: 210, A: 255}
			St = chart.Style{
					StrokeColor: ColorCyan,              
					FillColor:   ColorCyan.WithAlpha(64),
			}
		}

		All = append(All, chart.ContinuousSeries{
			Name:  Sc[i],
			Style: St,
			XValues: Axe_X,
			YValues: Axe_Y,
		})
	}
	DrawHisto(All, mat, name)
}

func DrawHisto(All []chart.Series, mat string,  name string) {

	graph := chart.Chart {
		Series: All,
	}
	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}
	f, _ := os.Create("graphs/" + name)
	defer f.Close()
	graph.Render(chart.PNG, f)
}