package main

/*
  Program that processes weather data information.
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/* WeatherData...  struct that holds weather data information */
type WeatherData struct {
	wban         int
	yearMonthDay int
	tMax         int
	tMin         int
	tAvg         int
	station      string
	location     string
}

/* Reads lines form a file and returns them in a string slice */
func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

/* Date conversion / month lookup formatter */
func dateConversion(date int) string {
	dateString := strconv.Itoa(date)
	dateStringYear := dateString[0:4]
	dateStringMonth := dateString[4:6]
	dateStringDay := dateString[6:8]
	if strings.Contains(dateStringMonth, "01") {
		dateStringMonth = "January"
	} else if strings.Contains(dateStringMonth, "02") {
		dateStringMonth = "February"
	} else if strings.Contains(dateStringMonth, "03") {
		dateStringMonth = "March"
	} else if strings.Contains(dateStringMonth, "04") {
		dateStringMonth = "April"
	} else if strings.Contains(dateStringMonth, "05") {
		dateStringMonth = "May"
	} else if strings.Contains(dateStringMonth, "06") {
		dateStringMonth = "June"
	} else if strings.Contains(dateStringMonth, "07") {
		dateStringMonth = "July"
	} else if strings.Contains(dateStringMonth, "08") {
		dateStringMonth = "August"
	} else if strings.Contains(dateStringMonth, "09") {
		dateStringMonth = "September"
	} else if strings.Contains(dateStringMonth, "10") {
		dateStringMonth = "October"
	} else if strings.Contains(dateStringMonth, "11") {
		dateStringMonth = "November"
	} else if strings.Contains(dateStringMonth, "12") {
		dateStringMonth = "December"
	}
	return dateStringMonth + " " + dateStringDay + " " + dateStringYear
}

/* Main routine that runs weather data program. */
func main() {
	var weatherInfo []WeatherData

	for true {
		fmt.Println("Please enter the name of the weather file to be processed: ")
		reader := bufio.NewReader(os.Stdin)
		fileName, _ := reader.ReadString('\n')
		fileName = strings.TrimSuffix(fileName, "\n")

		// Check if file exists and process file data
		if !strings.Contains(fileName, "tempData2015.txt") && !strings.Contains(fileName, "tempData2017.txt") {
			fmt.Println("Please enter a valid weather data file.")
		} else {
			lines, err := readLines(fileName)
			if err != nil {
				log.Fatalf("readLines: %s", err)
			}
			for i := range lines {
				splitLines := strings.Split(lines[i], " ")

				wban, wbanerr := strconv.Atoi(splitLines[0])
				if wbanerr != nil {
					fmt.Print(wbanerr)
				}
				yearMonthDay, yearMonthDayerr := strconv.Atoi(splitLines[1])
				if yearMonthDayerr != nil {
					fmt.Print(yearMonthDayerr)
				}
				tMax, tMaxerr := strconv.Atoi(splitLines[2])
				if tMaxerr != nil {
					fmt.Print(tMaxerr)
				}
				tMin, tMinerr := strconv.Atoi(splitLines[3])
				if tMinerr != nil {
					fmt.Print(tMinerr)
				}
				tAvg, tAvgerr := strconv.Atoi(splitLines[4])
				if tAvgerr != nil {
					fmt.Print(tAvgerr)
				}
				locationReplace := strings.Replace(splitLines[6], "_", " ", -1)
				weatherInfo = append(weatherInfo, WeatherData{wban, yearMonthDay, tMax, tMin, tAvg, splitLines[5], locationReplace})
			}

			// Use Variables
			var year int
			if strings.Contains(fileName, "tempData2015.txt") {
				year = 2015
			} else {
				year = 2017
			}
			location := ""
			var date int
			station := ""
			convertedDate := ""
			numberOfStations := len(weatherInfo) / 31

			// Find the maximum temperature
			fmt.Printf("\n1. What is the maximum temperature reported by any of the WBAN's during August %d?\n", year)
			maximum := 0

			for i := range weatherInfo {
				if weatherInfo[i].tMax > maximum {
					maximum = weatherInfo[i].tMax
					location = weatherInfo[i].location
					date = weatherInfo[i].yearMonthDay
					convertedDate = dateConversion(date)
					station = weatherInfo[i].station
				}
			}

			fmt.Printf("The max temperature recorded in August %d was %d on %s at %s in %s\n", year, maximum, convertedDate, location, station)

			// Find the minimum temperature
			fmt.Printf("\n2. What is the minimum temperature reported by any of the WBAN's during August %d?\n", year)
			minimum := 100000000

			for i := range weatherInfo {
				if weatherInfo[i].tMin < minimum {
					minimum = weatherInfo[i].tMin
					location = weatherInfo[i].location
					date = weatherInfo[i].yearMonthDay
					convertedDate = dateConversion(date)
					station = weatherInfo[i].station
				}
			}
			fmt.Printf("The min temperature recorded in August %d was %d on %s at %s in %s\n", year, minimum, convertedDate, location, station)

			// Find the average temperature
			fmt.Printf("\n3. What is the average for all 25 reporting stations in August %d?\n", year)
			var average int
			count := 0
			total := 0

			for i := range weatherInfo {
				count++
				total += weatherInfo[i].tAvg
			}

			average = total / count
			fmt.Printf("The average temperature recorded in August %d was %d.\n", year, average)

			// Find the hottest day
			fmt.Printf("\n4. What was the hottest day in Pennsylvania in August %d?\n", year)
			currentMaximumStation := 1
			totalTmax := 0
			tMaxAverage := 0
			var tMaxDate int
			tMaxStation := ""
			tMaxLocation := ""

			sort.Slice(weatherInfo, func(i int, j int) bool {
				return weatherInfo[i].yearMonthDay < weatherInfo[j].yearMonthDay
			})

			for i := range weatherInfo {
				totalTmax += weatherInfo[i].tMax
				if currentMaximumStation == numberOfStations {
					if (totalTmax / numberOfStations) > tMaxAverage {
						tMaxAverage = totalTmax / numberOfStations
						tMaxDate = weatherInfo[i].yearMonthDay
						convertedDate = dateConversion(tMaxDate)
						tMaxStation = weatherInfo[i].station
						tMaxLocation = weatherInfo[i].location
					}
					currentMaximumStation = 0
					totalTmax = 0
				}
				currentMaximumStation++
			}
			fmt.Printf("The hottest day in Pennsylvania in August %d was %d on %s at %s in %s.\n", year, tMaxAverage, convertedDate, tMaxLocation, tMaxStation)

			// Find the coldest day
			fmt.Printf("\n5. What was the coldest day in Pennsylvania in August %d?\n", year)
			currentMinimumStation := 1
			totalTmin := 0
			tMinAverage := 100000000
			var tMinDate int
			tMinStation := ""
			tMinLocation := ""

			for i := range weatherInfo {
				totalTmin += weatherInfo[i].tMin
				if currentMinimumStation == numberOfStations {
					if (totalTmin / numberOfStations) < tMinAverage {
						tMinAverage = totalTmin / numberOfStations
						tMinDate = weatherInfo[i].yearMonthDay
						convertedDate = dateConversion(tMinDate)
						tMinStation = weatherInfo[i].station
						tMinLocation = weatherInfo[i].location
					}
					currentMinimumStation = 0
					totalTmin = 0
				}
				currentMinimumStation++
			}
			fmt.Printf("The coldest day in Pennsylvania in August %d was %d on %s at %s in %s.\n", year, tMinAverage, convertedDate, tMinLocation, tMinStation)

			break
		}
	}
}
