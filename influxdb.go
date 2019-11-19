package main

import (
	"log"
	"time"

	client "https://github.com/influxdata/influxdb1-client/tree/master/v2"
)

const (
	db       = "f1"
	username = "admin"
	password = "admin"
	addr     = "localhost:8089"
)

// Point is a struct with data for sending to InfluxDB
type Point struct {
	tp *TelemetryPack
	t  time.Time
}

func influxDBSender(ch chan Point) {

	c, err := client.NewUDPClient(client.UDPConfig{
		Addr: addr,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	for p := range ch {
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database: db,
		})
		if err != nil {
			log.Fatal(err)
		}
		pt, err := client.NewPoint("TelemetryPack", nil, p.tp.ToMap(), p.t)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
		if err := c.Write(bp); err != nil {
			log.Fatal(err)
		}

	}

}
