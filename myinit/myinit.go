package myinit

import (
	"fmt"
)

var msg = ""

func init() {
	msg = "lovely message"
	fmt.Println("Done init")
}

func Show() {
	fmt.Println(msg)
}
