package file

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"Response"
	"types"
	"maps"
	"dataOp"
)

func ReadFile(Dat *types.Datas, File string, t int) (int) {

	csvfile, err := os.Open(File)
	if err != nil {
		Response.Print(fmt.Sprintf("%s\n", err))
		return (1)
	}
	r := csv.NewReader(csvfile)
	index := 0

	Dat.Table = make(map[int]types.Dat)
	Dat.Categ = make(map[int]string)
	if t == 1 {
		Dat.School = make(map[int]string)
		Dat.Mat = make(map[int]string)
	}
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
			if t == 1 {
				index = maps.Array_search(Dat.Categ, "Hogwarts House")
				if index == -1 {
					Response.Print("No Hause in your data")
					return (1)
				}
			}
		} else if line >= 1 {
			if t == 1 {
				RempHisto(record, Dat.School, Dat.Mat, index)
			}
			RempData(record, Dat.Table, Dat.Categ, 1)
		}
		line++
	}
	if line == 0 {
		return (1)
	}
	return (0)
}

func RempHisto(record []string, S map[int]string, M map[int]string, iS int) {

	if iS != -1 {
		if !dataOp.InArray(S, record[iS]) {
			S[len(S)] = record[iS]
		}
	}
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
			if IsNumeric(record[i]) {
				Data[i].Cat[len(Data[i].Cat)], _ = strconv.ParseFloat(record[i], 64)
			}
		}
	}
}

func IsNumeric(s string) (bool) {

    _, err := strconv.ParseFloat(s, 64)
    return (err == nil)
}
