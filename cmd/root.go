package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/paconte/torneito/model"
	"github.com/spf13/cobra"
)

var csvFile string
var games model.Games = model.Games{}
var tournaments []model.Tournament = []model.Tournament{}

var rootCmd = &cobra.Command{
	Use:   "torneito",
	Short: "Torneito is a very fast sport tournaments generator",
	Long: `A fast and flexible sport tournament generator built with
			love by paconte and friends in Go.
			Complete documentation is available at https://github.com/paconte/torneito`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("File to read: " + csvFile)
		records := Import(csvFile)
		games = games.AddFromRecords(records)
		tournaments = games.Tournaments()
		fmt.Println("Tournaments loaded: " + fmt.Sprintf("%d", len(tournaments)))
		fmt.Println("Games loaded: " + fmt.Sprintf("%d", len(games)))
		for {
			selectRootAction()
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

func selectRootAction() {
	prompt := promptui.Select{
		Label: "Select action",
		Items: []string{"Add", "Edit", "Delete", "List tournaments", "List games", "Export", "Exit"},
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
	case "Edit":
		EditGame()
	case "List tournaments":
		ListTournaments()
	case "List games":
		ListTournamentGames()
	case "Export":
		Export()
	case "Exit":
		fmt.Println("Bye!")
		os.Exit(1)
	}
}
