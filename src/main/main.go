package main

import (
	f "fmt"
	//_ "github.com/wdpress/gosample"
	. "strings"
	"gosample"
)

const (
	message  string = "hello, world."
	message2 string = "bye!"
	flg      bool   = true
)

func main() {
	//mes := "say yes."
	f.Println(ToUpper(gosample.Message))
	f.Println(flg)
	/*
	   if文テスト
	*/
	a, b := 100, 100
	if a > b {
		f.Println("a is larger than b")
	} else if a < b {
		f.Println("a is smaller than b")
	} else {
		f.Println("a equals b")
	}

}
