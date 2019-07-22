package main

import (
	"fmt"
)

func main() {
	buffered := make(chan string, 2)
	go func() {
		for {
			str, ok := <-buffered
			if !ok {
				fmt.Println(`no buffer`)
				return
			}
			fmt.Println(str)
		}
	}()

	buffered <- "1"
	buffered <- "2"
	// buffered <- "3"
	// buffered <- "4"
	close(buffered)
}
