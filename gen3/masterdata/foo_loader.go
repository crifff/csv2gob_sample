// +build !build_gob

package masterdata

import (
	"encoding/gob"
	"os"
	"time"
)

var FooList []*Foo

type Foo struct {
	Version   int
	Title     string
	ID        int
	BeginDate time.Time
}

func init() {
	file, err := os.Open("masterdata/Foo.gob")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	dec := gob.NewDecoder(file)
	if err := dec.Decode(&FooList); err != nil {
		panic(err)
	}
}
