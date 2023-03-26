package array_slice_test

import (
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int            //Declare new slice.
	t.Log(len(s0), cap(s0)) //Show length and capacity.
	s0 = append(s0, 1)      //Insert new data from tail.
	t.Log(len(s0), cap(s0)) //Show length and capacity.
	s1 := []int{1, 2, 3, 4} //Declare new slice and initialize.
	t.Log(len(s1), cap(s1)) //Show length and capacity.

	s2 := make([]int, 3, 5) //Another way to make new slice
	t.Log(len(s2), cap(s2)) //Show length and capacity.
	//t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
	t.Log(s2) //Showing all data.
	s2 = append(s2, 1)
	s2 = append(s2, 4)
	t.Log(s2)               //Showing all data after append.
	t.Log(len(s2), cap(s2)) //Show length and capacity.
}

func TestSliceGrowing(t *testing.T) { //Slice capacity doubles when there's not enough space.
	var growth []int
	for i := 1; i <= 16; i++ {
		growth = append(growth, i)
		t.Log(len(growth), cap(growth), growth)
	}
}

func TestSliceMemoryShare(t *testing.T) { //Slice interceptions share their memory and data.
	var months = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	var Q2 = months[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	Q2[2] = "July"
	t.Log(months)
}
