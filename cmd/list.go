package cmd

import (
	"fmt"

	"github.com/paconte/torneito/model"
)

func ListGames() {
	for i, g := range games {
		fmt.Printf("%d %s %s\n", i, g.Name, g.Serie)
	}
}

func ListTournaments() {
	for i, t := range tournaments {
		fmt.Printf("%d %s %s %s %s\n", i, t.Country, t.Serie, t.Division, t.Name)
	}
}

func ListTournamentGames() {
	var total int = 0
	t, err := readTournament()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, g := range games {
		result := ""
		for i, score := range g.Local {
			if i < len(g.Local)-1 {
				result = result + fmt.Sprintf("%d-%d / ", score, g.Visitor[i])
			} else {
				result = result + fmt.Sprintf("%d-%d", score, g.Visitor[i])
			}
		}

		local_team := g.Players[1] + " " + g.Players[0] + " && " + g.Players[3] + " " + g.Players[2]
		visitor_team := g.Players[5] + " " + g.Players[4] + " && " + g.Players[7] + " " + g.Players[6]

		if g.IsTournament(t) {
			fmt.Printf("%d %s %s %s %s %s\n", i, g.Round, g.Category, local_team, result, visitor_team)
			total++
		}
	}
	fmt.Printf("Total games: %d\n", total)
}

func readTournament() (model.Tournament, error) {
	var index int
	fmt.Print("Enter tournament index: ")
	fmt.Scanln(&index)
	if index < 0 || index >= len(tournaments) {
		return model.Tournament{}, fmt.Errorf("invalid tournament selection")
	}
	return tournaments[index], nil
}
