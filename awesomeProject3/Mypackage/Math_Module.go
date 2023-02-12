package Mypackage

import "fmt"

type Fraction struct {
	Num   int
	Denum int
}

func Addf(a, b Fraction) Fraction {
	C := Fraction{Num: (a.Num*b.Denum + a.Denum*b.Num), Denum: (b.Denum * a.Denum)}
	return C
}

func Substractf(a, b Fraction) Fraction {
	C := Fraction{Num: (a.Num*b.Denum - a.Denum*b.Num), Denum: (b.Denum * a.Denum)}
	return C
}

func Multiplyf(a, b Fraction) Fraction {
	C := Fraction{Num: (a.Num * b.Num), Denum: (b.Denum * a.Denum)}
	return C
}

func Dividef(a, b Fraction) Fraction {
	C := Fraction{Num: (a.Num * b.Denum), Denum: (a.Denum * b.Num)}
	return C
}

type Vector struct {
	X int
	Y int
}

func AddV(a, b Vector) Vector {
	C := Vector{X: (a.X + b.X), Y: (a.Y + b.Y)}
	return C
}

func SubtractV(a, b Vector) Vector {
	C := Vector{X: (a.X - b.X), Y: (a.Y - b.Y)}
	return C
}

func PrintSmth(a string) {
	fmt.Println(a)
}
