package dataOp

import (
	"types"
	"math"
	"fmt"
	"maths"
	"maps"
	"os"
	"log"
	"encoding/csv"
	"strings"
)

func GetMat(cat map[int]string, mat map[int]string) {

	for i := 6; i < len(cat); i++ {
		mat[len(mat)] = cat[i]
	}
}

func Calc(data types.Datas) (types.GraphHisto) {

	Save := types.GraphHisto{}
	Save.Table = make(map[int]types.Dat)
	Sc := [4]string{"Gryffindor", "Hufflepuff", "Ravenclaw", "Slytherin"}

	for i := 0; i < len(data.Mat); i++ {

		Dat := types.Dat{}
		Dat.Cat = make(map[int]float64)
		for z := 0; z < 4; z++ {
			Dat.Cat = Getdatas(Sc[z], i + 6, data.Table, data.School)
			Save.Table[len(Save.Table)] = Dat
		}
	}
	return (Save)
}

func Getdatas(search string, index int, data map[int]types.Dat, school map[int]string) (map[int]float64) {

	tmp := make(map[int]float64)

	for z := 0; z < len(data[0].Cat); z++ {
		if school[z] == search && !math.IsNaN(data[index].Cat[z]) {
			tmp[len(tmp)] = data[index].Cat[z]
		}
	}
	return (tmp)
}

func ScatterPlot(data types.Datas, Resp *types.SaveCor, verbose bool) {

	ReduceData(&data)
	for cols := 6; cols < len(data.Table); cols++ {

		for n_cols := 6; n_cols < len(data.Table); n_cols++ {

			if n_cols > cols {
				cor := maths.Correlation(data.Table[cols].Cat, data.Table[n_cols].Cat)
				if verbose == true {
					fmt.Printf("Correlation between %s | %s : %f\n", data.Categ[cols], data.Categ[n_cols], cor)
				}
				if math.Abs(cor) > Resp.Cor {
					Resp.Cor = math.Abs(cor)
					Resp.Name1 = data.Categ[cols]
					Resp.Name2 = data.Categ[n_cols]
					Resp.Map1 = data.Table[cols].Cat
					Resp.Map2 = data.Table[n_cols].Cat
				}
			}
		}
	}
}

func PairPlot(data types.Datas) (string, int) {

	file, err := os.Create("datasets/result.csv")
    if !checkError("Cannot create file", err) {
    	return "", 1
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    header := FormateHeader(data.Categ, 6)
    header = append(header, "School")
    Write(header, writer)
    ReduceData(&data)
    for z := 0; z < len(data.Table[0].Cat); z++{

    	ndat := ""
    	for i := 6; i < len(data.Table); i++ {

    		ndat += fmt.Sprintf("%f,", data.Table[i].Cat[z])
    	}
    	if data.School[z] != "" && ndat != "0.000000,0.000000,0.000000,0.000000,0.000000,0.000000,0.000000,0.000000,0.000000,0.000000,0.000000,0.000000,0.000000," {
    		Write(FormateData(ndat, data.School[z]), writer)
    	}
    }
    return "datasets/result.csv", 0
}

func Write(data []string, writer *csv.Writer) {

    err := writer.Write(data)
    if !checkError("Cannot write to file", err) {
    	return
    }
}

func checkError(message string, err error) (bool) {
    if err != nil {
        log.Fatal(message, err)
        return (false)
    }
    return (true)
}

func FormateHeader(data map[int]string, deb int) ([]string) {

	var tab []string

	for i := deb; i < len(data); i++ {
		tab = append(tab, data[i])
	}
	return (tab)
}

func FormateData(data string, sc string) ([]string) {

	var tab []string


	e := strings.Split(data, ",")
	for i := 0; i < len(e) - 1; i++ {
		tab = append(tab, e[i])
	}
	tab = append(tab, sc)
	return (tab)
}

func ReduceData(data *types.Datas) {

	for i := 6; i < len(data.Table); i++ {

		for z := 0; z < len(data.Table[i].Cat); z++ {
			if math.IsNaN(data.Table[i].Cat[z]) {
				DeleteData(data, z)
				delete(data.School, z)
			}
		}
	}
}

func DeleteData(data *types.Datas, z int) {

	for a := 6; a < len(data.Table); a++ {
		delete(data.Table[a].Cat, z)
		N := types.Dat{}
		N.Cat = maps.Reindex(data.Table[a].Cat)
		N.Cat = maps.Clean(N.Cat)
		data.Table[a] = N
	} 
}

func SupprUs(ids []int, data types.Datas) (types.Datas) {

	Dats := types.Datas{}
	Dats.Categ = data.Categ
	Dats.School = data.School
	Dats.Mat = data.Mat
	Dats.Table = make(map[int]types.Dat)

	for i := 0; i < len(data.Table); i++ {

		if !InIDs(ids, i) {
			Dats.Table[len(Dats.Table)] = data.Table[i]
		}
	}
	return (Dats)
}

func InIDs(ids []int, search int) (bool) {

	for i := 0; i < len(ids); i++ {
		if ids[i] == search {
			return (true)
		}
	}
	return (false)
}

func FormatData(res *types.DataTrain, data types.Datas) {

	res.Line = make(map[int]map[int]float64)

	for z := 0; z < len(data.Table[0].Cat); z++ {

		Adds := make(map[int]float64)
		for i := 0; i < len(data.Table); i++ {
			if math.IsNaN(data.Table[i].Cat[z]) {
				data.Table[i].Cat[z] = maths.Median(data.Table[i].Cat)
			}
			Adds[i] = data.Table[i].Cat[z]
		}
		res.Line[len(res.Line)] = Adds
	}
}

func RempY(sc string, data map[int]string) (map[int]float64) {

	res := make(map[int]float64)

	for i := 0; i < len(data); i++ {

		if sc == data[i] {
			res[len(res)] = 1
		} else {
			res[len(res)] = 0
		}
	} 
	return (res)
}