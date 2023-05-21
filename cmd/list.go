package cmd

import "fmt"

func ListGames() {
	for i, g := range games {
		fmt.Printf("%d %s (%s)\n", i, g.Name, g.Country)
	}
}
