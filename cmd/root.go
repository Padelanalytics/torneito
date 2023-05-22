package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/paconte/torneito/model"
	"github.com/spf13/cobra"
)

var csvFile string
var games model.Games = model.Games{}

var rootCmd = &cobra.Command{
	Use:   "torneito",
	Short: "Torneito is a very fast sport tournaments generator",
	Long: `A fast and flexible sport tournament generator built with
				  love by paconte and friends in Go.
				  Complete documentation is available at https://gohugo.io/documentation/`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("File to read: " + csvFile)
		records := readCsvFile(csvFile)
		games = games.FromRecords(records)
		for {
			selectAction()
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&csvFile, "csvFile", "f", "", "Path to csv file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func selectAction() {
	prompt := promptui.Select{
		Label: "Select action",
		Items: []string{"Add", "Delete", "List", "Export", "Exit"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Add":
		AddGame()
	case "Delete":
		RemoveGame()
	case "List":
		ListGames()
	case "Export":
		AddGame()
	case "Exit":
		fmt.Println("Bye!")
		os.Exit(1)
	}
}

func readCsvFile(filePath string) [][]string {
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
