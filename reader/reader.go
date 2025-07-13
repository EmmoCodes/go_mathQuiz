package reader

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

func ReadCSV() ([][]string, error) {

	fmt.Println("Please enter the filename with the calculations.")
	var file string = "questions.csv"
	// fmt.Scan(&file)

	f, err := os.Open(file)

	if err != nil {
		f.Close()
		return nil, errors.New("Failed to open file.")
	}

	csvReaver := csv.NewReader(f)
	records, err := csvReaver.ReadAll()
	if err != nil {
		f.Close()
		return nil, errors.New("Failed to read file.")
	}
	f.Close()
	return records, nil
}

func GetResults(userInput *string) (string, error) {
	records, err := ReadCSV()
	if err != nil {
		return "", errors.New("Failed to find file.")
	}

	var lastVal string
	for _, n := range records {
		fmt.Println(n)
		// take len - 1 to get the last element
		lastVal = n[len(n)-1]
		fmt.Println(lastVal)
	}
	return lastVal, nil
}

func CompareNumbers(userInput *string) {
	records, err := GetResults(userInput)
	if err != nil {
		return
	}
	fmt.Println(records)

}
