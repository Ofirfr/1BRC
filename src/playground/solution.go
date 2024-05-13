package playground

import (
	"1BRC/src/structs"
	"bufio"
	"log"
	"os"
)

func CalculateStatistics() map[string]*structs.CityResult {
	file, err := os.Open("../data/temperature_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	statistics := make(map[string]*structs.CityResult, 10000)
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go func() {
	// 	defer wg.Done()
	// 	scanner := bufio.NewScanner(file)
	// 	counter := 0
	// 	for scanner.Scan() {
	// 		if counter > 500000 {
	// 			continue
	// 		}
	// 		scanner.Text()
	// 		counter++
	// 	}
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	scanner := bufio.NewScanner(file)
	// 	counter := 0
	// 	for scanner.Scan() {
	// 		if counter < 500000 {
	// 			continue
	// 		}
	// 		scanner.Text()
	// 		counter++
	// 	}
	// }()
	// wg.Wait()
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		// if counter < 500000 {
		// 	continue
		// }
		scanner.Text()
		counter++
	}
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()

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
	// 		statistics[city] = &structs.CityResult{
	// 			Max:     -100,
	// 			Min:     100,
	// 			Sum:     0,
	// 			Count:   0,
	// 			Average: 0,
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
	// for _, cityStats := range statistics {
	// 	cityStats.Average = cityStats.Sum / cityStats.Count
	// }
	return statistics
}
