package array_slice_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int               //Length set but not initialized.
	arr1 := [4]int{1, 2, 3, 4}   //Length set and initialized.
	arr3 := [...]int{1, 3, 4, 5} //Length not defined by user, set to 4 for defined integers.
	arr1[1] = 5                  //set value for variable in array.
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTraverse(t *testing.T) { //Shown in recent codes.
	arr := [...]int{1, 1, 4, 5, 1, 4}
	for i, e := range arr {
		t.Log(i, e)
	}
}

func TestArrayIntercept(t *testing.T) { //Array interception, usage: [<front>:<rear(not included)>]
	arr := [...]int{1, 1, 4, 5, 1, 4}
	arrSec := arr[0:]
	arrSec1 := arr[1:5]
	t.Log(arrSec)
	t.Log(arrSec1)
}
