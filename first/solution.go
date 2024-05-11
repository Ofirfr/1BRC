package first

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const Max = 1
const Min = 2
const Sum = 3
const Count = 4

func CalculateStatistics() {
	file, err := os.Open("./data/temperature_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	statistics := make(map[string]map[uint8]float32)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Parse the line into city and temperature
		city_and_temp := strings.Split(line, ";")
		city := city_and_temp[0]
		temp_string := city_and_temp[1]
		temp, err := strconv.ParseFloat(temp_string, 32)
		if err != nil {
			log.Fatal(err)
		}

		if _, ok := statistics[city]; !ok {
			statistics[city] = make(map[uint8]float32)
			statistics[city][Max] = -100
			statistics[city][Min] = 100
		}

		// Calculate statistics
		statistics[city][Max] = max(statistics[city][Max], float32(temp))
		statistics[city][Min] = min(statistics[city][Min], float32(temp))
		statistics[city][Sum] = statistics[city][Sum] + float32(temp)
		statistics[city][Count] = statistics[city][Count] + 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
