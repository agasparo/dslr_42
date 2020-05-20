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
	"math"
	"io/ioutil"
)

func SaveFile(data map[int][]float64, name string) {

	var str string

	for i := 0; i < len(data); i++ {

		for a := 0; a < len(data[i]); a++ {
			str += fmt.Sprintf("%f", data[i][a])
			if a + 1 < len(data[i]) {
				str += ","
			}
		}
		str += "\n"
	}

	fd := []byte(str)
    err := ioutil.WriteFile(name, fd, 0644)
    check(err, name, 1)
}

func SaveHeader(data []string, name string) {

	var str string

	for a := 0; a < len(data); a++ {
		str += fmt.Sprintf("%s", data[a])
		if a + 1 < len(data) {
				str += ","
		}
	}
	str += "\n"
	fd := []byte(str)
    err := ioutil.WriteFile(name, fd, 0644)
    check(err, name, 1)
}

func SaveLines(data []string, name string) {

	var str string

	for a := 0; a < len(data); a++ {

		str += fmt.Sprintf("%s", data[a])
		if a + 1 < len(data) && a % 2 == 0 {
				str += ","
		}
		if a % 2 != 0 && a != 0 && a + 1 < len(data) {
			str += "\n"
		}
	}
	str += "\n"
	fd := []byte(str)
    err := ioutil.WriteFile(name, fd, 0644)
    check(err, name, 0)
}

func check(e error, name string, v int) {
    
    if e != nil {
        Response.Print(fmt.Sprintf("%s\n", e))
    } else {
    	if v == 1 {
    		Response.Sucess(fmt.Sprintf("File %s created", name))
    	}
    }
}

func ReadDat(name string) ([]float64) {

	var data, nul []float64

	csvfile, err := os.Open(name)
	if err != nil {
		Response.Print(fmt.Sprintf("%s\n", err))
		return (nul)
	}
	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			Response.Print(fmt.Sprintf("%s\n", err))
			return (nul)
		}
		for i := 0; i < len(record); i++ {
			conv, _ := strconv.ParseFloat(record[i], 64)
			data = append(data, conv)
		}
	}
	return (data)
}

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
			RempData(record, Dat.Table, Dat.Categ, 0, t)
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
			RempData(record, Dat.Table, Dat.Categ, 1, t)
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
		S[len(S)] = record[iS]
	}
}

func RempData(record []string, Data map[int]types.Dat, Categ map[int]string, t int, count int) {

	if t == 0 {

		for i := 0; i < len(record); i++ {
			n := types.Dat{}
			n.Cat = make(map[int]float64)
			Data[len(Data)] = n
			Categ[len(Categ)] = record[i]
		}
	} else {
		for i := 0; i < len(record); i++ {
			if IsNumeric(record[i]) && count == 0 {
				Data[i].Cat[len(Data[i].Cat)], _ = strconv.ParseFloat(record[i], 64)
			}
			if count == 1 {
				nb, err := strconv.ParseFloat(record[i], 64)
				if err != nil {
					Data[i].Cat[len(Data[i].Cat)] = math.NaN()
				} else {
					Data[i].Cat[len(Data[i].Cat)] = nb
				}
			}
		}
	}
}

func IsNumeric(s string) (bool) {

    _, err := strconv.ParseFloat(s, 64)
    return (err == nil)
}
