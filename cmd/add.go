package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
	"github.com/paconte/torneito/model"
)

func AddGame() model.Game {

	country := read(model.Countries(tournaments), "Enter country:")
	name := read(model.Names(tournaments), "Enter name:")
	serie := read(model.Series(tournaments), "Enter serie:")
	division := read(model.Divisions(tournaments), "Enter division:")
	date := read(games.Dates(), "Enter date:")
	round := read(games.Rounds(), "Enter round:")
	category := read(games.Categories(), "Enter category:")
	lastnames := games.Lastnames()
	firstnames := games.Firstnames()
	p1Last := read(lastnames, "Enter player 1 last name:")
	p1First := read(firstnames, "Enter player 1 first name:")
	p2Last := read(lastnames, "Enter player 2 last name:")
	p2First := read(firstnames, "Enter player 2 first name:")
	p3Last := read(lastnames, "Enter player 3 last name:")
	p3First := read(firstnames, "Enter player 3 first name:")
	p4Last := read(lastnames, "Enter player 4 last name:")
	p4First := read(firstnames, "Enter player 4 first name:")
	players := []string{p1Last, p1First, p2Last, p2First, p3Last, p3First, p4Last, p4First}
	scores := readScores()

	game := model.Game{
		Country:  country,
		Name:     name,
		Serie:    serie,
		Division: division,
		Date:     date,
		Teams:    2,
		Round:    round,
		Category: category,
		Players:  players,
		Sets:     2,
		Scores:   scores,
	}

	fmt.Println("Game:", game)
	return game
}

func read(completerContent []string, text string) string {
	completer := readline.NewPrefixCompleter(
		readline.PcItemDynamic(func(line string) []string { return completerContent }),
	)

	l, err := readline.NewEx(&readline.Config{AutoComplete: completer})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	l.CaptureExitSignal()

	fmt.Println(text)
	line, err := l.Readline()
	if err != nil {
		panic(err)
	}
	return line
}

func readScores() []int8 {
	var scores []int8

	l, err := readline.NewEx(&readline.Config{})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	l.CaptureExitSignal()

	iterate := true
	for iterate {
		fmt.Println("Enter scores:")
		line, err := l.Readline()
		if err != nil {
			panic(err)
		}
		words := strings.Fields(line)
		if len(words)%2 == 0 {
			for _, word := range words {
				s, err := strconv.Atoi(word)
				if err != nil {
					fmt.Printf("Invalid score %s", word)
					scores = []int8{}
					break
				}
				scores = append(scores, int8(s))
			}
			iterate = false
		}
	}

	return scores
}
