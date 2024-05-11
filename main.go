package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
    start := time.Now()

	calculate_statistics()

    elapsed := time.Since(start)
    log.Printf("Task took %s", elapsed)
}

func calculate_statistics() {
	file, err := os.Open("./data/temperature_data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	const Max = 1
	const Min = 2
	const Sum = 3
	const Count = 4
	statistics := make(map[string]map[uint8]float32)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        city_and_temp := strings.Split(scanner.Text(), ";")
		city := city_and_temp[0]
		temp_string := city_and_temp[1]
		temp, err := strconv.ParseFloat(temp_string, 32);
		if err != nil {
			log.Fatal(err)
		}

		if _, ok := statistics[city]; !ok {
			statistics[city] = make(map[uint8]float32)
			statistics[city][Max] = -100
			statistics[city][Min] = 100
		}
		statistics[city][Max] = max(statistics[city][Max], float32(temp))
		statistics[city][Min] = min(statistics[city][Min], float32(temp))
		statistics[city][Sum] = statistics[city][Sum] + float32(temp)
		statistics[city][Count]  = statistics[city][Count] + 1
    } 
		
	// for city := range statistics {
		// fmt.Println("City: ", city, " Max: ", statistics[city][Max], " Min: ", statistics[city][Min], " Average: ", statistics[city][Sum] / statistics[city][Count], "Count: ", statistics[city][Count])
	// }	

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}