package reader

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadCSV() ([][]string, error) {
	fmt.Println("Please enter the filename with the calculations.")
	var file string
	fmt.Scan(&file)

	f, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		defer f.Close()
		return nil, errors.New("Failed to open file.")
	}

	csvReaver := csv.NewReader(f)
	records, err := csvReaver.ReadAll()
	if err != nil {
		fmt.Println(err)
		defer f.Close()
		return nil, errors.New("Failed to read file.")
	}
	defer f.Close()
	return records, nil
}

func GetResults(userInput *string) error {
	winVal, loseVal, err := GetInputs(userInput)
	if err != nil {
		fmt.Println(err)
		return errors.New("Failed to get inputs.")
	}

	fmt.Printf("Right solved: %v\n", winVal)
	fmt.Printf("Not right: %v\n", loseVal)

	return nil
}

func GetInputs(userInput *string) (int, int, error) {
	var lastVal string
	var comparedVal int
	var winCounter, loseCounter int

	records, err := ReadCSV()
	if err != nil {
		fmt.Println(err)
		return 0, 0, errors.New("Failed to read in result")
	}

	fmt.Println("Please solve following calculation: ")
	for _, val := range records {
		fmt.Println(val[0])
		var input string
		fmt.Scanln(&input)

		// take len - 1 to get the last element
		lastVal = val[len(val)-1]

		userInputInt, err := ConvertToInt(input)
		if err != nil {
			return 0, 0, errors.New("Failed to convert string.")
		}

		lastValInt, err := ConvertToInt(lastVal)
		if err != nil {
			return 0, 0, errors.New("Failed to convert string.")
		}
		comparedVal = lastValInt - userInputInt

		if comparedVal == 0 {
			winCounter++
		} else {
			loseCounter++
		}

	}

	return winCounter, loseCounter, nil
}

func ConvertToInt(val string) (int, error) {
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return 0, errors.New("Failed to convert string into Int")
	}
	return valInt, nil
}
