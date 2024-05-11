package playground

import (
	"1BRC/src/structs"
	"bufio"
	"log"
	"os"
)

type CityStatistics struct {
	Max   float64
	Min   float64
	Sum   float64
	Count float64
}

func CalculateStatistics() map[string]structs.CityResult {
	file, err := os.Open("../data/temperature_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	statistics := make(map[string]*CityStatistics)

	scanner := bufio.NewScanner(file)
	lines := make([]string, 10)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	// 	// Parse the line into city and temperature
	// 	city_and_temp := strings.Split(line, ";")
	// 	city := city_and_temp[0]
	// 	temp_string := city_and_temp[1]
	// 	temp, err := strconv.ParseFloat(temp_string, 64)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	cityStats, ok := statistics[city]
	// 	if !ok {
	// 		statistics[city] = &CityStatistics{
	// 			Max:   -100,
	// 			Min:   100,
	// 			Sum:   0,
	// 			Count: 0,
	// 		}
	// 		cityStats = statistics[city]
	// 	}

	// 	cityStats.Max = max(cityStats.Max, temp)
	// 	cityStats.Min = min(cityStats.Min, temp)
	// 	cityStats.Sum = cityStats.Sum + temp
	// 	cityStats.Count = cityStats.Count + 1

	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	results := make(map[string]structs.CityResult, 10000)
	for city, cityStats := range statistics {
		results[city] = structs.CityResult{
			Max:     cityStats.Max,
			Min:     cityStats.Min,
			Average: cityStats.Sum / cityStats.Count,
		}
	}
	return results
}
