package first

import (
	"1BRC/src/structs"
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
	Count float64
}

func CalculateStatistics() map[string]structs.CityResult {
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
		temp, err := strconv.ParseFloat(temp_string, 64)
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
			cityStats = statistics[city]
		}

		statistics[city] = CityStatistics{
			Max:   max(cityStats.Max, temp),
			Min:   min(cityStats.Min, temp),
			Sum:   cityStats.Sum + temp,
			Count: cityStats.Count + 1,
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	results := make(map[string]structs.CityResult)
	for city, cityStats := range statistics {
		results[city] = structs.CityResult{
			Max:     cityStats.Max,
			Min:     cityStats.Min,
			Average: cityStats.Sum / cityStats.Count,
		}
	}
	return results
}
