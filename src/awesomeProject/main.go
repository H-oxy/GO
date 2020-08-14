package main

import (
	"fmt"

)

func main () {

fmt.Println("My multiplication: ",  multiply(5, 5))









}

func multiply(i  ...int)   int {

	mult := 1

	for _, v := range i {
		mult *= v


	}


	return mult + 1



}