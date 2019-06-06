package main

import (
	"csv2go/gen2/masterdata"
	"encoding/gob"
	"os"
)

var MasterList = map[string]interface{}{
	"Foo": masterdata.FooList,
}

func main() {
	for i := range MasterList {
		file, err := os.Create("masterdata/" + i + ".gob")
		if err != nil {
			panic(err)
		}
		enc := gob.NewEncoder(file)
		if err := enc.Encode(MasterList[i]); err != nil {
			panic(err)
		}
	}
}
