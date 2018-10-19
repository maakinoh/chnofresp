package lib

import (
	"fmt"
)

type Buzz struct {
	Prt
}

//Create a new Buzz
func NewBuzz(nextPr Printer) Buzz {
	return Buzz{Prt: NewPrt(nextPr)}
}

func (b Buzz) Print(num int64) {
	if num%5 == 0 {
		fmt.Println("Buzz")
	}
}
