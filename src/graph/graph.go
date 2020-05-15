package graph

import (
	"types"
	"maths"
	"sort"
	"os"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
	"github.com/mattn/go-pairplot"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"log"
	"Response"
	"fmt"
	"images"
)

func DrawOne(Data types.GraphHisto, Sc [4]string, mat string, g int, verbose bool) {

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

		if verbose == true {

			str := fmt.Sprintf("-------------------------------------\n")
			str += fmt.Sprintf("Subject : %s\n", mat)
			str += fmt.Sprintf("School : %s\n\n", Sc[i])
			str += fmt.Sprintf("Count : %f\n", C)
			str += fmt.Sprintf("Average : %f\n", M)
			str += fmt.Sprintf("Standard deviation : %f\n", ET)
			str += fmt.Sprintf("-------------------------------------\n")

			Response.PrintVerbose(str)
		}

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
	DrawHisto(All, mat)
	images.Resize("graphs/" + mat + ".png")
}

func DrawHisto(All []chart.Series, mat string) {

	graph := chart.Chart {
		XAxis: chart.XAxis{
			Name: "Ecart type",
		},
		YAxis: chart.YAxis{
			Name: "nombres de notes",
		},
		Series: All,
	}
	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}
	f, _ := os.Create("graphs/" + mat + ".png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func DrawScatterPlot(Res types.SaveCor) {

	viridisByY := func(xr, yr chart.Range, index int, x, y float64) drawing.Color {
		return chart.Viridis(y, yr.GetMin(), yr.GetMax())
	}

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeWidth:      chart.Disabled,
					DotWidth:         5,
					DotColorProvider: viridisByY,
				},
				XValues: maths.MaptoSlice(Res.Map1),
				YValues: maths.MaptoSlice(Res.Map2),
			},
		},
	}
	f, _ := os.Create("graphs/ScatterPlot.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func DrawPairPlot(name string) {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	pp, err := pairplot.NewPairPlotCSV(name)
	if err != nil {
		log.Fatal(err)
	}
	pp.SetHue("School")
	p.HideAxes()
	p.Add(pp)
	p.Save(8*vg.Inch, 8*vg.Inch, "graphs/pair_plot.png")
}