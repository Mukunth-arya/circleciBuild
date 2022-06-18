package main

import (
	"fmt"
	"testing"
)

func Testsample(test *testing.T) {

	var sample = 20
	var datas = 30

	if datas != sample {
		test.Errorf("The value is not equal")

	} else if datas == sample {
		fmt.Println("The values is equal ?!!! ")
	}

}
