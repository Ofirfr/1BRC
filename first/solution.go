package first

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type CityStatistics struct {
	Max   float64
	Min   float64
	Sum   float64
	Count int64
}

func CalculateStatistics() {
	file, err := os.Open("./data/temperature_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	statistics := make(map[string]CityStatistics)

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

		cityStats, ok := statistics[city]
		if !ok {
			statistics[city] = CityStatistics{
				Max:   -100,
				Min:   100,
				Sum:   0,
				Count: 0,
			}
		}

		cityStats = CityStatistics{
			Max:   max(cityStats.Max, float64(temp)),
			Min:   min(cityStats.Min, float64(temp)),
			Sum:   cityStats.Sum + float64(temp),
			Count: cityStats.Count + 1,
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
