package main

import (
	"fmt"
	"os"

	"github.com/nats-io/nats"
)

// We use globals because it's a small application demonstrating NATS.

var nc *nats.Conn

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments. Need NATS server address.")
		return
	}
	var err error

	nc, err = nats.Connect(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	nc.QueueSubscribe("TimeTeller", "TimeTellers", replyWithTime)
	select {} // Block forever
}
