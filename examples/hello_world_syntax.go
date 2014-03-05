package main

import (
	"fmt"
	"strings"
)

//START OMIT
func main() {
	fmt.Println(getHellos(5) + "world")
}

func getHellos(numHellos int) string {
	return strings.Repeat("Hello ", numHellos)
}

//END OMIT
