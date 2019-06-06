// +build build_gob

package main

import (
	"csv2go/gen3/masterdata"
	"encoding/gob"
	"io/ioutil"
	"os"
)

var MasterList = map[string]interface{}{
	"Foo": &masterdata.FooList,
}

var LoaderList = map[string]func([]byte) error{
	"Foo": masterdata.LoadFooList,
}

func main() {
	for i := range LoaderList {
		raw, err := ioutil.ReadFile("masterdata/" + i + ".json")
		if err != nil {
			panic(err)
		}
		if err := LoaderList[i](raw); err != nil {
			panic(err)
		}

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
