package main

import (
	"csv2go/gen1/masterdata"
	"fmt"
)

func main() {
	for i := range masterdata.FooList {
		fmt.Printf("FooList[%d] is %+v\n", i, masterdata.FooList[i])
	}
}
