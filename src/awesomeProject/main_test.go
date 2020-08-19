package main

import "testing"






func TestDivi(t *testing.T) {

	type test struct {
		fnumber int
		snumber int
		answer  int
	}

	tests := []test{
		test{25, 5, 5},
		test{36, 6, 7},
	}

	for _, v := range tests {
		x := divi(v.fnumber, v.snumber)
        if x != v.answer {

        	t.Error("Expected ", 5, "Got", x)

		}
	}

}