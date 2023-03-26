package condition_loop_test

import "testing"

func TestIfCondition(t *testing.T) { //If...else usage, similar to other languages.
	if a := 1 == 1; a {
		t.Log("true")
	} else {
		t.Log("false")
	}
}

func TestSwitchCondition(t *testing.T) { //Switch usage, similar to other languages.
	for i := 1; i <= 12; i++ {
		switch i {
		case 1, 3, 5, 7, 9:
			t.Log("odd")
		case 2, 4, 6, 8, 10:
			t.Log("even")
		default:
			t.Log("too high")
		}
	}
}
