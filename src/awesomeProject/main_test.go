package main

import "testing"






func TestMultiply(t *testing.T){

	type test struct {
		data []int
		answer int
	}

	tests :=[]test{
		  test{[]int{5, 5}, 25},
		  test{[]int{6, 6}, 36},
		  test{[]int{10, 10},100},

	}

  for _, v := range tests{

		x := multiply(v.data...)
		if x != v.answer {

		t.Error("it is suposed to be:", v.answer, "GOT", x )



	}
	}







}