package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/nats-io/stan.go"
)

func logCloser(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("close error: %s", err)
	}
}

func run() error {
	conn, err := stan.Connect(
		"test-cluster",
		"test-client",
		stan.NatsURL("nats://localhost:4222"),
	)
	if err != nil {
		return err
	}
	defer logCloser(conn)

	wg := &sync.WaitGroup{}

	sub, err := conn.Subscribe("counter", func(msg *stan.Msg) {
		// Print the value and whether it was redelivered.
		fmt.Printf("seq = %d [redelivered = %v]\n", msg.Sequence, msg.Redelivered)

		// Add jitter..
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		log.Print(msg)
		// Mark it is done.
		wg.Done()
	})
	if err != nil {
		return err
	}
	defer logCloser(sub)

	// Publish up to 10.
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// var d byte[] =
		buf := []byte("exexeeexeywiu3y8yi3x")
		err := conn.Publish("counter", buf)
		if err != nil {
			return err
		}
	}

	// Wait until all messages have been processed.
	wg.Wait()

	return nil
}
