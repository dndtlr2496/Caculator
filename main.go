package main

import (
	"fmt"

	"github.com/dndtlr2496/Caculator/module"
)

func main() {
	fmt.Println(module.Add(3, 4))
	fmt.Println(module.Minus(4, 2))
	fmt.Println(module.Multiply(6, 3))
	fmt.Println(module.Division(6, 3))
}
