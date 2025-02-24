package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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

func InsertMatch(matchData *MatchCreate) (*MatchDataOut, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	fmt.Println(id)

	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := createdAt

	query := `INSERT INTO match (
		id, hash, map, team_a_name, team_b_name, team_a_score, team_b_score,
		team_a_score_first_half, team_b_score_first_half, 
		team_a_score_second_half, team_b_score_second_half,
		team_a_overtime_rounds_won, team_b_overtime_rounds_won,
		created_at, updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = DB.Exec(
		query,
		id, matchData.FileHash, matchData.Map,
		matchData.TeamAName, matchData.TeamBName, matchData.TeamAScore, matchData.TeamBScore,
		matchData.TeamAScoreFirstHalf, matchData.TeamBScoreFirstHalf,
		matchData.TeamAScoreSecondHalf, matchData.TeamBScoreSecondHalf,
		matchData.TeamAOvertimeRoundsWon, matchData.TeamBOvertimeRoundsWon,
		createdAt, updatedAt,
	)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// now insert player data
	for steamID, playerData := range matchData.PlayerData {
		query := `INSERT INTO player_match (
			match_id, steam_id, player_name, final_team, kills, deaths, diff, kpr, dpr, adr,
			pct_rounds_with_mk, opening_kills_per_round, win_pct_after_opening_kill, impact, kast, rating
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

		_, err = DB.Exec(
			query,
			id, steamID, playerData.Name, playerData.FinalTeam, playerData.Kills, playerData.Deaths,
			playerData.Diff, playerData.KPR, playerData.DPR, playerData.ADR,
			playerData.PctRoundsWithMK, playerData.OpeningKillsPerRound, playerData.WinPctAfterOpeningKill,
			playerData.Impact, playerData.KAST, playerData.Rating,
		)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	return &MatchDataOut{ID: id}, nil
}

func GetMatchById(id string) (*MatchData, error) {
	query := `SELECT * FROM matches WHERE id = ?`
	row := DB.QueryRow(query, id)

	var matchData MatchData
	var matchId string
	err := row.Scan(
		&matchId, &matchData.FileHash, &matchData.Map,
		&matchData.TeamAName, &matchData.TeamBName, &matchData.TeamAScore, &matchData.TeamBScore,
		&matchData.TeamAScoreFirstHalf, &matchData.TeamBScoreFirstHalf,
		&matchData.TeamAScoreSecondHalf, &matchData.TeamBScoreSecondHalf,
		&matchData.TeamAOvertimeRoundsWon, &matchData.TeamBOvertimeRoundsWon,
		&matchData.CreatedAt, &matchData.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	query = `SELECT * FROM players WHERE match_id = ?`
	rows, err := DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	matchData.PlayerData = make(map[string]Player)
	for rows.Next() {
		var playerData Player
		var steamID string
		err := rows.Scan(
			&matchId, &steamID, &playerData.Name, &playerData.FinalTeam, &playerData.Kills, &playerData.Deaths,
			&playerData.Diff, &playerData.KPR, &playerData.DPR, &playerData.ADR,
			&playerData.PctRoundsWithMK, &playerData.OpeningKillsPerRound, &playerData.WinPctAfterOpeningKill,
			&playerData.Impact, &playerData.KAST, &playerData.Rating,
		)

		if err != nil {
			return nil, err
		}

		matchData.PlayerData[steamID] = playerData
	}

	return &matchData, nil
}

func GetMatchIdFromFileHash(fileHash string) (string, error) {
	query := `SELECT id FROM matches WHERE hash = ?`
	row := DB.QueryRow(query, fileHash)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}
