package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/rs/xid"
)

func main() {
	var (
		count    int
		timeStr  string
		decode   string
		machine  bool
		pid      bool
		counter  bool
		timeInfo bool
	)

	flag.IntVar(&count, "n", 1, "Number of IDs to generate")
	flag.StringVar(&timeStr, "t", "", "Generate ID with specific time (format: 2006-01-02T15:04:05)")
	flag.StringVar(&decode, "d", "", "Decode XID and show its internal parts")
	flag.BoolVar(&machine, "m", false, "Show machine ID part when decoding")
	flag.BoolVar(&pid, "p", false, "Show process ID part when decoding")
	flag.BoolVar(&counter, "c", false, "Show counter part when decoding")
	flag.BoolVar(&timeInfo, "time", false, "Show timestamp part when decoding")
	flag.Parse()

	// Decode mode
	if decode != "" {
		id, err := xid.FromString(decode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding XID: %v\n", err)
			os.Exit(1)
		}

		if !machine && !pid && !counter && !timeInfo {
			// If no specific parts are requested, show all
			machine = true
			pid = true
			counter = true
			timeInfo = true
		}

		if timeInfo {
			fmt.Printf("Timestamp: %v\n", id.Time())
		}
		if machine {
			fmt.Printf("Machine: %x\n", id.Machine())
		}
		if pid {
			fmt.Printf("PID: %d\n", id.Pid())
		}
		if counter {
			fmt.Printf("Counter: %d\n", id.Counter())
		}
		return
	}

	// Generate mode
	var t time.Time
	if timeStr != "" {
		var err error
		t, err = time.Parse("2006-01-02T15:04:05", timeStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing time: %v\n", err)
			os.Exit(1)
		}
	}

	for i := 0; i < count; i++ {
		var id xid.ID
		if timeStr != "" {
			id = xid.NewWithTime(t)
		} else {
			id = xid.New()
		}
		fmt.Println(id.String())
	}
}
