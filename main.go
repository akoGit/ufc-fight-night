package main

import (
    "encoding/json"
	"fmt"
	"io"
    "os"
    "math/rand"
    "time"
    "github.com/fatih/color"
)

type Fight struct {
    Promotion  string `json:"promotion"`
    Event      string `json:"event"`
    Date       string `json:"date"`
    Fighter_01 string `json:"fighter_1"` 
    Fighter_02 string `json:"fighter_2"` 
}


func main() {
	file, err := os.Open("Desktop/fight/res.json")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var fights []Fight

	err = json.Unmarshal(content, &fights)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	now := time.Now()
	currentYear := now.Year()

	var eligibleFights []Fight

	for _, fight := range fights {
		fightDate, err := time.Parse("2006-01-02", fight.Date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}

		if fightDate.Year() == currentYear {
			eligibleFights = append(eligibleFights, fight)
		}
	}

	if len(eligibleFights) > 0 {
		randomIndex := rand.Intn(len(eligibleFights))
		selectedFight := eligibleFights[randomIndex]

		promotion, event, fighter1, fighter2, date :=
			selectedFight.Promotion, selectedFight.Event, selectedFight.Fighter_01, selectedFight.Fighter_02, selectedFight.Date

		color.Red("| %s %s | %s vs %s | Date: %s |\n", promotion, event, fighter1, fighter2, date)
	} else {
		fmt.Println("No fights found in the current year.")
	}
}
