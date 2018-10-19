package main

import (
	"fmt"

	"github.com/maakinoh/chnofresp/lib"
)

func main() {
	var f = lib.NewFizz(lib.NewBuzz(nil))
	fmt.Println("Bitte Zahl eingeben")

	f.Print(15)

}
