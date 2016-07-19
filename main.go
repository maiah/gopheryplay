package main

import (
	"fmt"
	"github.com/maiah/gopheryplay/consumer"
)

func main() {
	c := make(chan string)

	go func() {
		c <- "A gopher playing here"
	}()

	m := <-c
	fmt.Println(m)
	speak()
	speak2()
	consumer.Speak()
}
