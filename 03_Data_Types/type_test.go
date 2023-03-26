package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) { //Implicit type conversion
	var (
		a       = 1
		b int64 = 2
		c MyInt
	)
	b = int64(a)
	c = MyInt(a)
	t.Log(a, b, c)
}

func TestPointer(t *testing.T) { //Pointer type demo
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) { //String printing demo
	var s string
	s = "The Go programming language"
	t.Log("#" + s + "#")
	t.Log(len(s))
}
