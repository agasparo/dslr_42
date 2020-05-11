package norm

import (
	"maths"
	"types"
)

/*func Normalize(Data types.Datas, kilometrage float64) (float64) {

	return ((kilometrage - maths.Min(Data.Kilometre)) / (maths.Max(Data.Kilometre) - maths.Min(Data.Kilometre)))
}

func Denormalize(Data types.Datas, prix float64) (float64) {

	return ((prix * (maths.Max(Data.Price) - maths.Min(Data.Price))) + maths.Min(Data.Price))
}*/

func NormalizeAllData(Data *types.Datas) {

	for i := 0; i < len(Data.Table); i++ {

		minK := maths.Min(Data.Table[i].Cat)
		maxK := maths.Max(Data.Table[i].Cat)

		for j := 0; j < len(Data.Table[i].Cat); j++ {
			Data.Table[i].Cat[j] = (Data.Table[i].Cat[j] - minK) / (maxK - minK)
		}
	}
}