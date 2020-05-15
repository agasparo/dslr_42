package say

import (
	"Response"
	"input"
	"fmt"
)

func Histogram() {

	for {
		resp, _ := input.ReadSTDIN("Quel cours de Poudlard a une répartition des notes homogènes entre les quatres maisons ?", 0)
		if int(resp) == 5 {
			Response.Sucess("Good job")
			return
		} else {
			Response.Print("retry")
		}
	}
}

func PairPlot() {
	fmt.Println("À partir de cette visualisation, quelles caractéristiques allez-vous utiliser pour entraîner votre prochaine régression logistique ?")
}