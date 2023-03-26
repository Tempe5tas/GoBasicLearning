package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b) //Not equal.
	//t.Log(a == c)	//Compile error, lengths not equal.
	t.Log(a == d)
}

const (
	read  int = 4
	write int = 2
	exec  int = 1
)

func TestBitClear(t *testing.T) {
	var a = 7
	a = a &^ write //Cleaning write flag.
	t.Log(a, a&read == read, a&write == write, a&exec == exec)
}
