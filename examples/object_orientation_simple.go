package main

import (
	"fmt"
)

//START OMIT
type littleBox struct {
	toads            int
	numBaseballCards int
}

func main() {
	myLittleBox := littleBox{2, 7}
	fmt.Printf("My little box has %d baseballcards and %d toad(s)\n", myLittleBox.numBaseballCards, myLittleBox.toads)
}

//END OMIT
