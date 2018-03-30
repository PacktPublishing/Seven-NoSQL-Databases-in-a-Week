package main

import (
	"log"
	"time"
	"github.com/influxdata/influxdb/client/v2"
)

const (
	database  = "testdb"
	username = "root"
	password = "root"
)

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
	// Create a point and add to batch
	tags := map[string]string{}
	fields1 := map[string]interface{}{
		"http": "GET /cgi-bin/try/ HTTP/1.0",
		"status": 200,
		"duration": 3395,
		"ip": "192.168.2.20",
	}

	pt1, err := client.NewPoint("apachelog_go", tags, fields1, time.Unix(0, 1511361120000000000))
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt1)
	fields2 := map[string]interface{}{
		"http": "GET / HTTP/1.0",
		"status": 200,
		"duration": 2216,
		"ip": "127.0.0.1",
	}
	pt2, err := client.NewPoint("apachelog_go", tags, fields2, time.Unix(0, 1511361180000000000))
	bp.AddPoint(pt2)
	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
}