package main

import (
	"fmt"

	"github.com/residenti/go_module_sample/calc"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(calc.Sum(1, 1))
}
