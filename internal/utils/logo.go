package utils

import (
	_ "embed"
	"fmt"
	"math"
)

//go:embed logo.txt
var logo string

func PrintLogo() {

	logo := string(logo)
	for j := 0; j < len(logo); j++ {
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, logo[j])
	}
	fmt.Println()
}

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}
