package nhl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Team struct {
	TeamKey string
}

type Roster struct {
	Forwards   []Player `json:"forwards"`
	Defensemen []Player `json:"defensemen"`
	Goalies    []Player `json:"goalies"`
}

type Player struct {
	Headshot  string `json:"headshot"`
	FirstName struct {
		Default string `json:"default"`
	} `json:"firstName"`
	LastName struct {
		Default string `json:"default"`
	} `json:"lastName"`
	SweaterNumber int `json:"sweaterNumber"`
}

func (t Team) List() Roster {
	var roster Roster

	err := json.Unmarshal([]byte(t.fetchRoster()), &roster)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}

	return roster
}

func (t Team) fetchRoster() string {
	url := fmt.Sprintf("https://api-web.nhle.com/v1/roster/%s/20242025", t.TeamKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return ""
	}

	return string(body)
}
