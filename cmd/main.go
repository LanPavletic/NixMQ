package main

import (
	"crypto/tls"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/LanPavletic/nixMQ"
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

	broker := nixmq.New()

	cert, err := tls.LoadX509KeyPair("/home/lan/faks/diploma/nixMQ/certs/server.crt", "/home/lan/faks/diploma/nixMQ/certs/server.key")
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	tcp := listeners.NewTCP("127.0.0.1", "1883")
	tls := listeners.NewTLS("127.0.0.1", "8883", tlsConfig)
	broker.AddListener(tcp)
	broker.AddListener(tls)
	broker.Start()

	<-done
}
