package main

import (
	"os"
	"os/signal"
	"syscall"

	nixmq "github.com/LanPavletic/nixMQ"
	"github.com/LanPavletic/nixMQ/listeners"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		done <- true
	}()

	b := nixmq.New()

	l := listeners.NewTCP("localhost", "1883")
	b.AddListener(l)

	b.Start()
	<-done
}
