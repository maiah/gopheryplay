package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	go func() {
		c <- "A gopher playing here"
	}()

	m := <-c
	fmt.Println(m)
}
