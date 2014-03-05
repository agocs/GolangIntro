package main

import (
	"fmt"
)

//START OMIT
type littleBox struct {
	toads            int
	numBaseballCards int
}

type bigBox struct {
	little   littleBox
	numBooks int
}

func (bb bigBox) makeString() string {
	return fmt.Sprintf("contains %d books and %d baseball cards", bb.numBooks, bb.little.numBaseballCards)
}

func main() {
	myLittleBox := littleBox{2, 7}
	myBigBox := bigBox{myLittleBox, 12}
	fmt.Println(myBigBox.makeString())
}

//END OMIT
