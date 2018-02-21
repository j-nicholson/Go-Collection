package main

/* Program to alter a list of floating point numbers */
import (
	"fmt"
	"strconv"
	"unicode"
)

/* Function that checks to see if a slice contains a number */
func contains(s []float64, item float64) bool {
	for _, i := range s {
		if i == item {
			return true
		}
	}
	return false
}

/* Function that checks if a string contains a number */
func containsNum(s string) bool {
	for _, i := range s {
		if unicode.IsNumber(i) {
			return true
		}
	}
	return false
}

func main() {
	end := false
	var floatingSlice []float64
	var command string
	var item string
	fmt.Println("* * * Floating-point program started * * *")

	for !end {
		fmt.Print("\nEnter a command: ")
		fmt.Scanf("%s %s", &command, &item)

		switch command {
		// Insert command option
		case "Insert":
			if containsNum(item) {
				item, err := strconv.ParseFloat(item, 64)
				floatingSlice = append(floatingSlice, item)
				fmt.Printf("\nThe array currently contains: \n")
				for i := 0; i < len(floatingSlice); i++ {
					fmt.Printf("Values[%d] = %.5f\n", i, floatingSlice[i])
				}
				fmt.Print(err)
			} else {
				fmt.Print("\nInvalid option: please choose [Insert, Delete, Sum, End]")
			}
			// Delete command option
		case "Delete":
			if containsNum(item) {
				item, err := strconv.ParseFloat(item, 64)
				if contains(floatingSlice, item) {
					fmt.Printf("\nThe array currently contains: \n")
					for i := 0; i < len(floatingSlice); i++ {
						if item == floatingSlice[i] {
							floatingSlice = append(floatingSlice[:i], floatingSlice[i+1:]...)
						}
					}
					for i := range floatingSlice {
						fmt.Printf("Values[%d] = %.5f\n", i, floatingSlice[i])
					}
					fmt.Print(err)
				} else {
					fmt.Printf("\nThat number was not in the list.\n")
				}
			} else {
				fmt.Print("\nInvalid option: please choose [Insert, Delete, Sum, End]")
			}
			// Sum command option
		case "Sum":
			var sum float64
			for i := 0; i < len(floatingSlice); i++ {
				sum += floatingSlice[i]
			}
			fmt.Printf("\nThe total is %.1f\n", sum)
			// End command option
		case "End":
			end = true
			fmt.Print("\n* * * Floating point program ended * * *\n")
			// Default option
		default:
			fmt.Print("\nInvalid option: please choose [Insert, Delete, Sum, End]")
		}
	}
}
