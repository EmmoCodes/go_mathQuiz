package reader

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func ReadCSV() ([][]string, error) {
	fmt.Println("Please enter the filename with the calculations.")
	var file string
	fmt.Scan(&file)

	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file.\nerr: %v", err)
	}

	csvReaver := csv.NewReader(f)
	records, err := csvReaver.ReadAll()
	if err != nil {
		defer f.Close()
		return nil, fmt.Errorf("Failed to read file.\nerr: %v", err)
	}
	defer f.Close()
	return records, nil
}

func GetResults(userInput *string) error {
	var wg sync.WaitGroup

	winVal, loseVal := GetInputs(userInput, &wg)

	fmt.Printf("Right solved: %v\n", winVal)
	fmt.Printf("Not right: %v\n", loseVal)

	switch loseVal {
	case 0:
		fmt.Println("BRAIN!")
	case 1:
		fmt.Println("AWESOME!")
	case 2:
		fmt.Println("GOOD!")
	default:
		fmt.Println("You did good! Keep learning.")
	}

	return nil
}

func GetInputs(userInput *string, wg *sync.WaitGroup) (int, int) {
	var lastVal, userTimer string
	var comparedVal int
	var winCounter, loseCounter int

	records, err := ReadCSV()
	if err != nil {
		return 0, 0
	}

	fmt.Println("How long your timer should be?")
	fmt.Scanln(&userTimer)
	timerToInt, err := ConvertToInt(userTimer)
	if err != nil {
		return 0, 0
	}

	fmt.Printf("Thanks. Your timer is set to: %.0v sec.\n", timerToInt)

	fmt.Println("Please solve following calculation: ")

	timeout := time.After(time.Duration(timerToInt) * time.Second)
	wg.Add(1)
	go func() {
		for _, val := range records {
			fmt.Println(val[0])
			var input string
			fmt.Scanln(&input)

			// take len - 1 to get the last element
			lastVal = val[len(val)-1]
			if len(lastVal) > 0 {

				userInputInt, err := ConvertToInt(input)
				if err != nil {
					return
				}

				lastValInt, err := ConvertToInt(lastVal)
				if err != nil {
					return
				}
				comparedVal = lastValInt - userInputInt

				if comparedVal == 0 {
					winCounter++
				} else {
					loseCounter++
				}
			} else {
				return
			}
		}
	}()

	<-timeout
	fmt.Println("Time's up!")
	defer wg.Done()
	return winCounter, loseCounter
}

func ConvertToInt(val string) (int, error) {
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("Failed to convert to integer.\nerr: %v", err)
	}
	return valInt, nil
}
