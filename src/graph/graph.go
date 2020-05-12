package graph

import (
	"types"
	//"github.com/wcharczuk/go-chart"
	//"os"
	"fmt"
	//"maths"
	"dataOp"
	//"maps"
)

func DrawOne(Data types.GraphHisto, Sc [4]string, name string, mat string, i int) {

	/*All := make(map[int]int)

	for i := 0; i < 4; i++ {
		min := maths.Min(Data.Stats)
		key := maps.ArraySearchFloat(Data.Stats, min)
		i = key
		Data.Stats[key] = maths.Max(Data.Stats) + Data.Stats[key]
	}
	fmt.Println(All)*/
	DrawHisto(Data, Sc, name, mat, i)
}

func DrawHisto(Data types.GraphHisto, Sc [4]string, name string, mat string, i int) {
	
	v := make(map[int]float64)
	var n float64

	v[0], n = dataOp.CountSearch(Data.Table[i + 0].Cat, 4, 1, 0.0)
	v[1], n = dataOp.CountSearch(Data.Table[i + 0].Cat, 2, 1, n)
	v[2], n = dataOp.CountSearch(Data.Table[i + 0].Cat, 4, 3, n)
	v[3], _ = dataOp.CountSearch(Data.Table[i + 0].Cat, 1, 1, n)

	v[4], n = dataOp.CountSearch(Data.Table[i + 1].Cat, 4, 1, 0.0)
	v[5], n = dataOp.CountSearch(Data.Table[i + 1].Cat, 2, 1, n)
	v[6], n = dataOp.CountSearch(Data.Table[i + 1].Cat, 4, 3, n)
	v[7], _ = dataOp.CountSearch(Data.Table[i + 1].Cat, 1, 1, n)

	v[8], n = dataOp.CountSearch(Data.Table[i + 2].Cat, 4, 1, 0.0)
	v[9], n = dataOp.CountSearch(Data.Table[i + 2].Cat, 2, 1, n)
	v[10], n = dataOp.CountSearch(Data.Table[i + 2].Cat, 4, 3, n)
	v[11], _ = dataOp.CountSearch(Data.Table[i + 2].Cat, 1, 1, n)

	v[12], n = dataOp.CountSearch(Data.Table[i + 3].Cat, 4, 1, 0.0)
	v[13], n = dataOp.CountSearch(Data.Table[i + 3].Cat, 2, 1, n)
	v[14], n = dataOp.CountSearch(Data.Table[i + 3].Cat, 4, 3, n)
	v[15], _ = dataOp.CountSearch(Data.Table[i + 3].Cat, 1, 1, n)

	fmt.Printf("%s%s%s\n", "Element", "Value", "Histogram")

    for z := 0; z < len(v); z++ {
        fmt.Printf("%d %f        ", z, v[z])

        for j := 1.0; j < v[z]; j++ {
            fmt.Printf("%c", '∎')
        }
        fmt.Println()
    }
}