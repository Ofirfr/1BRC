package main

import (
	"1BRC/src/first"
	"1BRC/src/structs"
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	log.Println("First solution: ", measureTime(first.CalculateStatistics))
	// log.Println("Second solution: ", measureTime(second.CalculateStatistics))
}

func measureTime(f func() map[string]structs.CityResult) time.Duration {
	startTime := time.Now()

	f()

	endTime := time.Now()

	duration := endTime.Sub(startTime)

	return duration
}
