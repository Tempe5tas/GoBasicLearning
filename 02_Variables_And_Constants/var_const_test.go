package var_const_test

import (
	"testing"
)

func TestFibArray(t *testing.T) {
	var ( //Variables SHOULD be used in program.
		a = 1
		b = 1
	)
	t.Log(a) //Testing log, returning Fibonacci array.
	for i := 1; i < 10; i++ {
		t.Log(b)
		a, b = b, a+b //Difference from languages such as C++.
	}
}

const ( //Unused constants may cause warnings.
	Mon int = iota + 1
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

func TestDate(t *testing.T) {
	t.Log(Mon, Tue, Wed)
}

const (
	Execute = 1 << iota //Bitwise operation
	Write
	Read
)

func TestFilePermission(t *testing.T) {
	var a = 3
	t.Log("Read\t", a&Read == Read) //Bitwise operation
	t.Log("Write\t", a&Write == Write)
	t.Log("Execute\t", a&Execute == Execute)
}
