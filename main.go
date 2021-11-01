package main

import (
	"fmt"

	h "github.com/thaibui2308/api-wrapper/helpers"
)

func main() {
	initialArray, pageNumber := h.MapInitials("877")
	from, to := h.SearchInterval("Kraft", initialArray)
	fmt.Println(initialArray, pageNumber)
	fmt.Println(from, to)
}
