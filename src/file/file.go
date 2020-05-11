package file

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"Response"
	"types"
)

func ReadFile(Dat *types.Datas, File string) (int) {

	csvfile, err := os.Open(File)
	if err != nil {
		Response.Print(fmt.Sprintf("%s\n", err))
		return (1)
	}
	r := csv.NewReader(csvfile)

	Dat.Table = make(map[int]types.Dat)
	Dat.Categ = make(map[int]string)
	line := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			Response.Print(fmt.Sprintf("%s\n", err))
			return (2)
		}
		if line == 0 {
			RempData(record, Dat.Table, Dat.Categ, 0)
		} else if line >= 1 {
			RempData(record, Dat.Table, Dat.Categ, 1)
		}
		line++
	}
	if line == 0 {
		return (1)
	}
	return (0)
}

func RempData(record []string, Data map[int]types.Dat, Categ map[int]string, t int) {

	if t == 0 {

		for i := 0; i < len(record); i++ {
			n := types.Dat{}
			n.Cat = make(map[int]float64)
			Data[len(Data)] = n
			Categ[len(Categ)] = record[i]
		}
	} else {
		for i := 0; i < len(record); i++ {
			Data[i].Cat[len(Data[i].Cat)], _ = strconv.ParseFloat(record[i], 64)
		}
	}
}