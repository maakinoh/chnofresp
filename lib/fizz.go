package lib

import (
	"fmt"
)

type Fizz struct {
	Prt
}

//
func NewFizz(prnt Printer) Fizz {
	return Fizz{Prt: NewPrt(prnt)}
}

func (f Fizz) Print(num int64) {
	if num%3 == 0 {
		fmt.Println("Fizz")
	} else {
		f.nextPrinter.Print(num)
	}

}
