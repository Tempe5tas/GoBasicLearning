package function_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func MultiValue() (int, int) {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10), rand.Intn(100)
}

func VerySlowFunction(op int) int {
	i := 0
	rand.Seed(time.Now().Unix())
	for rand.Intn(100) != op {
		time.Sleep(1 * time.Millisecond)
		i++
	}
	return i
}

func spentTime(inner func(op int) int) func(testInt int) {
	return func(n int) {
		start := time.Now()                                          //Record start time.
		inner(n)                                                     //Execute inner function.
		fmt.Println("Time spent:", time.Since(start).Milliseconds()) //Calculate spent time.
	}
}

func IntSum(nums ...int) int { //multi parameters
	ans := 0
	for _, num := range nums {
		ans += num
	}
	return ans
}

func CleanOutput() {
	fmt.Println("Cleaning resources... Please wait")
}

func TestDefer(t *testing.T) {
	defer CleanOutput() //defer is always reachable after function execution.
	fmt.Println("Function start.")
	panic("Function execution failed.")
	fmt.Println("Unreachable printing.") //Println after panic, unreachable.
}

func TestFunctions(t *testing.T) {
	t.Log(MultiValue())             //Return multiple value
	spentTime(VerySlowFunction)(67) //Test function execution time.
	t.Log(IntSum(1, 2, 3, 4))       //Test summary function.
}
