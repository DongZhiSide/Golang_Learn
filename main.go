package main

import (
	"fmt"
	"math/big"
)

func main() {
	var z1 big.Int
	z1.SetUint64(123)                     // z1 := 123
	z2 := new(big.Rat).SetFloat64(0.9999) // z2 := 5/4
	z3 := new(big.Float).SetInt(&z1)      // z3 := 123.0
	fmt.Println(z1, z2, z3)
	f1 := &big.Float{}
	f1.SetString("0.999")
	f2 := &big.Float{}
	f2.SetString("0.001")
	f1.Add(f1, f2)
	fmt.Println(f1.String())
}
