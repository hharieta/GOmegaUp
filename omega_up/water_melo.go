package main

import (
	"fmt"
)

func main() {
	var size_melon int
	fmt.Scan(&size_melon)

	if size_melon > 2 && size_melon%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
