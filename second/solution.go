package second

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

const Max = 1
const Min = 2
const Sum = 3
const Count = 4
const NumOfParsers = 1
const NumOfAggregators = 1

func CalculateStatistics() {
	var wgParser sync.WaitGroup
	var wgAggregator sync.WaitGroup

	lines := make(chan string)
	dataPoints := make(chan DataPoint)
	result := make(chan map[string]map[uint8]float32)
	go producer(lines)
	for i := 0; i < NumOfParsers; i++ {
		wgParser.Add(1)
		go func() {
			defer wgParser.Done()
			log.Println("Starting parser")
			parser(lines, dataPoints)
		}()
	}
	for i := 0; i < NumOfAggregators; i++ {
		wgAggregator.Add(1)
		go func() {
			defer wgAggregator.Done()
			log.Println("Starting aggregator")
			aggregator(dataPoints, result)
		}()
	}
	go func() {
		wgParser.Wait()
		close(dataPoints)
	}()
	go func() {
		wgAggregator.Wait()
		close(result)
	}()

	for _ = range result {
		log.Println("Hey")
	}
}

func producer(c chan string) {
defer close(c)
	file, err := os.Open("./data/temperature_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		c <- line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type DataPoint struct {
	city string
	temp float32
}

func parser(c chan string, d chan DataPoint) {
	for line := range c {
		// Parse the line into city and temperature
		city_and_temp := strings.Split(line, ";")
		city := city_and_temp[0]
		temp_string := city_and_temp[1]
		temp, err := strconv.ParseFloat(temp_string, 32)
		if err != nil {
			log.Fatal(err)
		}
		d <- DataPoint{city, float32(temp)}
	}
}

func aggregator(d chan DataPoint, r chan map[string]map[uint8]float32) {
	statistics := make(map[string]map[uint8]float32)

	for dataPoint := range d {
		city, temp := dataPoint.city, dataPoint.temp
		if _, ok := statistics[city]; !ok {
			statistics[city] = make(map[uint8]float32)
			statistics[city][Max] = -100
			statistics[city][Min] = 100
		}
		statistics[city][Max] = max(statistics[city][Max], float32(temp))
		statistics[city][Min] = min(statistics[city][Min], float32(temp))
		statistics[city][Sum] = statistics[city][Sum] + float32(temp)
		statistics[city][Count] = statistics[city][Count] + 1
	}

	r <- statistics
}