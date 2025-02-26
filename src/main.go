package main

import (
	"fmt"
	"github.com/brettmcginnis/nhl-slack-icon-pack/src/download"
	"github.com/brettmcginnis/nhl-slack-icon-pack/src/nhl"
	"github.com/mrz1836/go-sanitize"
	"os"
	"path"
	"strings"
)

func main() {
	saveRoster([]nhl.Team{
		{TeamKey: "SEA"},
		{TeamKey: "CAR"},
		{TeamKey: "BOS"},
		{TeamKey: "CBJ"},
		{TeamKey: "BUF"},
		{TeamKey: "NJD"},
		{TeamKey: "DET"},
		{TeamKey: "NYI"},
		{TeamKey: "FLA"},
		{TeamKey: "NYR"},
		{TeamKey: "MTL"},
		{TeamKey: "PHI"},
		{TeamKey: "OTT"},
		{TeamKey: "PIT"},
		{TeamKey: "TBL"},
		{TeamKey: "WSH"},
		{TeamKey: "TOR"},
		{TeamKey: "CHI"},
		{TeamKey: "ANA"},
		{TeamKey: "COL"},
		{TeamKey: "CGY"},
		{TeamKey: "DAL"},
		{TeamKey: "EDM"},
		{TeamKey: "MIN"},
		{TeamKey: "LAK"},
		{TeamKey: "NSH"},
		{TeamKey: "SJS"},
		{TeamKey: "STL"},
		{TeamKey: "UTA"},
		{TeamKey: "VAN"},
		{TeamKey: "WPG"},
		{TeamKey: "VGK"},
	})
}

func saveRoster(teams []nhl.Team) {
	for _, team := range teams {
		roster := team.List()

		saveImages(team.TeamKey, roster.Forwards)
		saveImages(team.TeamKey, roster.Defensemen)
		saveImages(team.TeamKey, roster.Goalies)
	}
}

func saveImages(teamName string, players []nhl.Player) {
	dir, _ := os.Getwd()

	for _, player := range players {

		fileName := fmt.Sprintf("%d_%s_%s.png", player.SweaterNumber, sanitize.AlphaNumeric(player.FirstName.Default, false), sanitize.AlphaNumeric(player.LastName.Default, false))
		directory := path.Join(dir, "images", teamName)
		err := os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}

		err = download.Image(player.Headshot, path.Join(directory, strings.ToLower(fileName)))
		if err != nil {
			fmt.Println(err)
		}
	}
}
