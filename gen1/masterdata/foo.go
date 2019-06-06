package masterdata

import (
	"csv2go/util"
	"time"
)

type Foo struct{
	Version   int
	Title     string
	ID        int
	BeginDate time.Time
}

var FooList = []*Foo{
	{
		Version: 12345,
		Title: "hoge",
		ID: 1,
		BeginDate: util.ParseDateTime("2019-01-01T00:00:00+09:00"),
	},
	{
		Version: 12345,
		Title: "fuga",
		ID: 2,
		BeginDate: util.ParseDateTime("2019-01-02T00:00:00+09:00"),
	},
}