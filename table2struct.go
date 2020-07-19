package main

import (
"fmt"
"github.com/gohouse/converter"
)

func main() {
	err := converter.NewTable2Struct().
		SavePath(`/Users/${whoami}/go/src/go-daydaystudy/model/model.go`).
		Dsn("user:pwd@tcp(127.0.0.1:3306)/base?charset=utf8").
		Run()
	fmt.Println(err)
}
