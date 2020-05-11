package maps

import (

)

func Array_search(array map[int]string, to_search string) (int) {

	for i := 0; i < len(array); i++ {

		if strings.Index(array[i], to_search) != -1 {
			return (i)
		}
	}
	return (-1)
}

func Reindex(data map[int]string) (map[int]string) {

	tab := make(map[int]string)	

	for i := getminkey(data); i < len(data); i++ {

		if data[i] != "" {
			tab[len(tab)] = data[i]
		}
	}

	if getminkey(data) == len(data) {
		tab[0] = data[len(data)]
	}
	return (tab)
}

func getminkey(data map[int]string) (int) {

	min := -1

	for index, element := range data {

		if element != "" && (min == -1 || index < min) {
			min = index
		}
	}
	return (min)
}

func Clean(data map[int]string) (map[int]string) {

	tab := make(map[int]string)

	for i := 0; i < len(data); i++ {

		if data[i] != "" {
			tab[len(tab)] = data[i]
		}
	}
	return (tab)
}