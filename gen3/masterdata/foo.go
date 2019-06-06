// +build build_gob

package masterdata

import (
	"encoding/json"
	"time"
)

type Foo struct {
	Version   int
	Title     string
	ID        int
	BeginDate time.Time
}

var FooList = []*Foo{}

func LoadFooList(raw []byte) error {
	if err := json.Unmarshal(raw, &FooList); err != nil {
		return err
	}
	return nil
}

func GetFooList() interface{} {
	return FooList
}
