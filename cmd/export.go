package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)

// Import imports the games from a csv file
// The file must be a csv file with ; as separator and # as comment
func Import(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	csvReader.Comment = '#'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

// Export exports the games to a csv file
// The file name is the same as the input file with a timestamp
func Export() {
	f := csvFile + "." + time.Now().Format("20060102150405")

	file, err := os.Create(f)
	if err != nil {
		log.Fatal("Cannot create file "+f, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	for _, value := range games {
		if err := writer.Write(value.ToRecord()); err != nil {
			log.Fatal("Cannot write games into "+f, err)
		}
	}
}
