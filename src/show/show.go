package show

import (
	"fmt"
	"text/tabwriter"
	"os"
)

func Header(Data map[int]string) (string) {

	str := "\t"

	for i := 0; i < len(Data); i++ {
		str += Data[i]
		if i  + 1 < len(Data) {
			str += "\t"
		}
	}
	return (str)
}

func FormatLine(str string, tab map[int]string) {

	ntab := make(map[int]string)
	ntab[0] = str
	for i := 0; i < len(tab); i++{
		ntab[len(ntab)] = tab[i]
	}
	doTable(ntab)
}

func doTable(tab map[int]string) {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	for i := 0; i < len(tab); i++{
		fmt.Fprintln(w, tab[i])
	}
    fmt.Fprintln(w)
    w.Flush()
}