package main


import (
	"fmt"
	"encoding/csv"
	"os"
)


func main() {
	// opening csv file

	file , err := os.Open("matches.csv")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close() 

	reader := csv.NewReader(file)

	// Reading headrs

	headers, err := reader.Read()

	if err != nil {
		fmt.Println("Error reading header:", err)
		return
	}

	var columnIndex int = -1

	for i, header := range headers {
		if header == "winner" {
			columnIndex = i
			break
		}
	}

	if columnIndex == -1 {
		fmt.Println("Column 'winner' not found in CSV.")
		return
	}

	fmt.Println("Winner Column:")
	winners := make(map[string]int)
	
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		if record[columnIndex] != "" {
			winners[record[columnIndex]]++
		}
	}

	fmt.Println("Totla wins by Each Team:")

	for team, count := range winners {
		fmt.Printf("%s: %d wins\n", team, count)
	}
}
