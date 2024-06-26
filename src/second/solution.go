package second

import (
	"1BRC/src/structs"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

const NumOfRows = 1000000

const Max = 1
const Min = 2
const Sum = 3
const Count = 4
const NumOfProducers = 1
const NumOfParsers = 1
const NumOfAggregators = 1
const LinesChannelBuffer = 1000000
const DataChannelBuffer = 1000000
const ResultChannelBuffer = 1000000

func CalculateStatistics() map[string]*structs.CityResult {
	var wgParser sync.WaitGroup
	var wgAggregator sync.WaitGroup
	var wgProducer sync.WaitGroup

	lines := make(chan string, LinesChannelBuffer)
	dataPoints := make(chan DataPoint, DataChannelBuffer)
	result := make(chan map[string]*structs.CityResult, ResultChannelBuffer)

	for i := 0; i < NumOfProducers; i++ {
		wgProducer.Add(1)
		start := NumOfRows / NumOfProducers * i
		end := NumOfRows / NumOfProducers * (i + 1)
		go func() {
			defer wgProducer.Done()
			log.Println("Starting producer")
			producer(lines, start, end)
		}()
	}
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
		wgProducer.Wait()
		close(lines)
	}()
	go func() {
		wgParser.Wait()
		close(dataPoints)
	}()
	go func() {
		wgAggregator.Wait()
		close(result)
	}()

	return <-result
}

func producer(c chan string, start int, end int) {
	file, err := os.Open("../data/temperature_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if count >= start && count < end {
			c <- line
		}
		count += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type DataPoint struct {
	city string
	temp float64
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
		d <- DataPoint{city, temp}
	}
}

func aggregator(d chan DataPoint, r chan map[string]*structs.CityResult) {
	statistics := make(map[string]*structs.CityResult, 10000)

	for dataPoint := range d {
		city, temp := dataPoint.city, dataPoint.temp
		cityResult, ok := statistics[city]
		if !ok {
			emptyCityResult := &structs.CityResult{
				Max:     -100,
				Min:     100,
				Sum:     0,
				Count:   0,
				Average: 0,
			}
			statistics[city] = emptyCityResult
			cityResult = emptyCityResult
		}
		cityResult.Max = max(cityResult.Max, temp)
		cityResult.Min = min(cityResult.Min, temp)
		cityResult.Sum = cityResult.Sum + temp
		cityResult.Count = cityResult.Sum + 1
	}

	r <- statistics
}
