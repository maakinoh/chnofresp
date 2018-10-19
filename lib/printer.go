package lib

import (
	"fmt"
)

type Printer interface {
	Print(num int64)
}

type Prt struct {
	nextPrinter Printer
}

func NewPrt(next Printer) Prt {
	p := Prt{next}
	return p
}

func (p *Prt) Print() {
	fmt.Println("Not Implementet")
}
