package Response

import (
	"github.com/fatih/color"
)

func Print(str string) {

	color.Red(str)
}

func Sucess(str string) {
	
	color.Green(str)
}