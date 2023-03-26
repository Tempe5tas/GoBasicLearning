package condition_loop_test

import "testing"

func TestLoop(t *testing.T) { //Basic for statement usage.
	for i := 1; i <= 5; i++ {
		t.Log(i, i%2 == 0)
	}
}

func TestForRange(t *testing.T) { //For...range usage, array traversal.
	ans, nums := 0, [...]int{1, 2, 3, 4, 5}
	for _, num := range nums {
		ans += num
	}
	t.Log(ans)
}
