package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("Starting the server")
	laddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:20777")
	if err != nil {
		log.Fatal(err)
	}

	con, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	fmt.Println("Server started")
	buf := make([]byte, 1289)

	ch := make(chan Point, 1000)
	for i := 0; i < 5; i++ {
		go influxDBSender(ch)
	}

	for {
		_, err := con.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		tp, err := NewTelemetryPack(buf)
		if err != nil {
			log.Fatal(err)
		}
		p := Point{tp: tp, t: time.Now()}
		ch <- p

	}
}
