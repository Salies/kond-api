package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	//"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GetPlayersFromSteam(playersData *SteamPlayers) (map[string]interface{}, error) {
	steamIDs := strings.Join(playersData.Players, ",")
	url := fmt.Sprintf(
		"https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002?key=%s&steamids=%s",
		SteamApiKey, steamIDs,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func InsertMatch(matchData *MatchCreate) error {
	id, err := gonanoid.New()
	if err != nil {
		return err
	}

	fmt.Println(id)

	//createdAt := time.Now().Format(time.RFC3339)
	//updatedAt := createdAt

	return nil
}
