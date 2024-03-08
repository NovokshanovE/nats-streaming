package main

import (
	"fmt"
	"os"
	"time"
)

func closeCLReader(done chan bool) {
	done <- true
}
func CLReader(done chan bool) {
	defer closeCLReader(done)
	defer closeCLReader(done)
	// defer closeCLReader(done)
	var read string
	time.Sleep(time.Duration(100 * time.Millisecond))
	// fmt.Print("Input \"exit\" to exit CLI \nAdd order ID:")
	// fmt.Fscan(os.Stdin, &read)
	for read != "exit" {

		fmt.Print("Input \"exit\" to exit CLI or ID if you want to see order info\nInput:")
		fmt.Fscan(os.Stdin, &read)
		if read == "exit" {

			continue
		}

		fmt.Printf("%s\n", cache.to_json(read))
	}

}
