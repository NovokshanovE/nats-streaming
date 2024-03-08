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

		fmt.Print("Commands:\nexit - shut down\nlist - list of IDs\ninput ID if you want to see order info\nInput:")
		fmt.Fscan(os.Stdin, &read)
		if read == "exit" {

			continue
		}
		if _, ok := cache.Orders[read]; ok {
			fmt.Printf("%s\n", cache.to_json(read))

			// if val == nil

			// fmt.Print(val)
		} else if read == "list" {
			for key := range cache.Orders {
				fmt.Printf("ID: %s\n", key)
			}
		} else {
			err := fmt.Errorf("ERROR: try to get order by ID = %s", read)
			fmt.Errorf("\n%e\n", err)
			fmt.Printf("ERROR: try to get order by ID = %s(there is no such order)\n", read)
		}

	}

}
