package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1416
	var tt int
	var figura int
	At := 0.00

	fmt.Scan(&tt)

	for tt >= 1 {
		fmt.Scan(&figura)
		switch figura {
		case 1:
			var l1 float64
			fmt.Scan(&l1)
			var area float64 = math.Pow(l1, 2)
			At += area
		case 2:
			var lado1, lado2 float64
			fmt.Scan(&lado1, &lado2)
			var area_r float64 = lado1 * lado2
			At += area_r
		case 3:
			var r float64
			fmt.Scan(&r)
			var area_c float64 = PI * math.Pow(r, 2)
			At += area_c
		case 4:
			var b, h float64
			fmt.Scan(&b, &h)
			var area_tr float64 = (b * h) / 2
			At += area_tr
		}

		tt--
	}

	AT := fmt.Sprintf("%.2f", At)
	fmt.Println(AT)
}
