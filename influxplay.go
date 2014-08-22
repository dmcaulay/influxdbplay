package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/influxdb/influxdb/client"
)

func main() {
	c, err := client.New(&client.ClientConfig{Database: "test_db"})
	if err != nil {
		log.Fatal(err)
	}

	// first run
	// series := []*client.Series{
	// 	&client.Series{
	// 		Name:    "log_lines",
	// 		Columns: []string{"line"},
	// 		Points:  [][]interface{}{{"here's some useful log info from paul@influx.com"}},
	// 	},
	// }

	// second run
	// series := []*client.Series{
	// 	&client.Series{
	// 		Name:    "response_times",
	// 		Columns: []string{"code", "value", "controller_action"},
	// 		Points:  [][]interface{}{{200, 234, "users#show"}},
	// 	},
	// 	&client.Series{
	// 		Name:    "user_events",
	// 		Columns: []string{"type", "url_base", "user_id"},
	// 		Points:  [][]interface{}{{"add_friend", "friends#show", 23}},
	// 	},
	// 	&client.Series{
	// 		Name:    "cpu_idle",
	// 		Columns: []string{"value", "host"},
	// 		Points:  [][]interface{}{{88.2, "serverA"}},
	// 	},
	// }

	// js example in go
	duration := 24 * time.Hour
	startTime := time.Now().Add(-duration)
	eventTypes := []string{"click", "view", "post", "comment"}

	cpuSeries := &client.Series{
		Name:    "cpu_idle",
		Columns: []string{"time", "value", "hostName"},
		Points:  [][]interface{}{},
	}

	eventSeries := &client.Series{
		Name:    "customer_events",
		Columns: []string{"time", "customerId", "type"},
		Points:  [][]interface{}{},
	}

	r := rand.New(rand.NewSource(startTime.UnixNano()))
	for i := time.Duration(0); i < duration; i += time.Minute {
		currentTimeMs := startTime.Add(i).UnixNano() / int64(time.Millisecond)

		hostName := fmt.Sprintf("server%d", r.Intn(100))
		value := r.Intn(100)
		pointValues := []interface{}{currentTimeMs, value, hostName}
		cpuSeries.Points = append(cpuSeries.Points, pointValues)

		for j := 0; j < r.Intn(10); j++ {
			customerId := r.Intn(1000)
			eventTypeIndex := r.Intn(4)
			eventValues := []interface{}{currentTimeMs, customerId, eventTypes[eventTypeIndex]}
			eventSeries.Points = append(eventSeries.Points, eventValues)
		}
	}
	series := []*client.Series{cpuSeries, eventSeries}

	// data, err := json.MarshalIndent(cpuSeries, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(string(data))

	err = c.WriteSeries(series)
	if err != nil {
		log.Fatal(err)
	}
}
