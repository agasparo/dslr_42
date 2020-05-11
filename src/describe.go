package main

import (
	"file"
	"fmt"
	"types"
	"show"
	"maths"
	"math"
	"os"
	"Response"
	"strings"
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
	res := file.ReadFile(&Data, args[0])
	if res != 0 {
		return
	}
	head := show.Header(Data.Categ)
	show.FormatLine(head, Describe(Data.Table))
}

func Describe(data map[int]types.Dat) (map[int]string) {

	var count, mean, std, min, max, per25, per50, per75 float64
	Tab := make(map[int]string)
	Tab[0] = "Count\t"
	Tab[1] = "Mean\t"
	Tab[2] = "std\t"
	Tab[3] = "Min\t"
	Tab[4] = "25%\t"
	Tab[5] = "50%\t"
	Tab[6] = "75%\t"
	Tab[7] = "Max\t"

	for i := 0; i < len(data); i++ {

		if !math.IsNaN(data[i].Cat[0]) {

			count = maths.Count(data[i].Cat)
			Tab[0] += fmt.Sprintf("%f", count)

			mean = maths.Mean(count, data[i].Cat)
			Tab[1] += fmt.Sprintf("%f", mean)

			std = maths.Std(mean, count, data[i].Cat)
			Tab[2] += fmt.Sprintf("%f", std)

			min = maths.Min(data[i].Cat)
			Tab[3] += fmt.Sprintf("%f", min)

			per25 = maths.Percentile(count, 4, data[i].Cat)
			Tab[4] += fmt.Sprintf("%f", per25)

			per50 = maths.Percentile(count, 2, data[i].Cat)
			Tab[5] += fmt.Sprintf("%f", per50)

			per75 = maths.Percentile( 3 * count, 4, data[i].Cat)
			Tab[6] += fmt.Sprintf("%f", per75)

			max = maths.Max(data[i].Cat)
			Tab[7] += fmt.Sprintf("%f", max)
		}

		if i + 1 < len(data) {
			AddTab(Tab)
		}
	}
	return (Tab)
}

func AddTab(dat map[int]string) {

	for i := 0; i < len(dat); i++ {

		dat[i] += "\t"
	}
}