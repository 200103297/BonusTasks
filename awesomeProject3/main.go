package main

import (
	"awesomeProject3/Mypackage"
	"fmt"
)

func main() {
	ar := Mypackage.Fraction{2, 3}
	br := Mypackage.Fraction{1, 3}
	fmt.Println(Mypackage.Addf(ar, br))
	vr := Mypackage.Vector{2, 2}
	se := Mypackage.Vector{5, 5}
	fmt.Println(Mypackage.AddV(vr, se))
}
