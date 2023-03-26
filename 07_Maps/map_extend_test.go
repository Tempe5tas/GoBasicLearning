package map_test

import "testing"

func TestMapWithFunction(t *testing.T) { //Map implements factory mode.
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapToSet(t *testing.T) { //Map implements Set collection.
	m := map[int]bool{}
	m[1] = true
	t.Log(m[1], m[3], len(m))
	delete(m, 1)
	m[3], m[5] = true, true
	t.Log(m[1], m[3], len(m))
}
