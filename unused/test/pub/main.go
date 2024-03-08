package main

import (
	"io"
	"log"

	stan "github.com/nats-io/go-nats-streaming"
)

// Convenience function to log the error on a deferred close.
func logCloser(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("close error: %s", err)
	}
}

func main() {
	// Specify the cluster (of one node) and some client id for this connection.
	conn, err := stan.Connect("test-cluster", "test-client")
	if err != nil {
		log.Print(err)
		return
	}
	defer logCloser(conn)

	// Now the patterns..
}
