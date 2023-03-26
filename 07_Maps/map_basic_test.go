package map_test

import "testing"

func TestMapInit(t *testing.T) { //Different ways to declare maps.
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2], len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Log(m2, len(m2))
	m3 := make(map[int]int, 10)
	t.Log(len(m3))
}

func TestAccessKey(t *testing.T) { //Check if a key is valid.
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	if v, ok := m1[3]; ok {
		t.Log(v)
	} else {
		t.Log("Nil")
	}
}

func TestMapRange(t *testing.T) { //For...range usage for maps.
	m1 := map[string]string{"a": "aa", "b": "bb"}
	for a, b := range m1 {
		t.Log(a, b)
	}
	t.Log(m1)
}
