package api

type (
	SteamPlayers struct {
		Players []string `json:"players"`
	}

	Player struct {
		FinalTeam              uint8   `json:"final_team"`
		Name                   string  `json:"name"`
		Kills                  uint32  `json:"kills"`
		Deaths                 uint32  `json:"deaths"`
		Diff                   int32   `json:"diff"`
		KPR                    float32 `json:"kpr"`
		DPR                    float32 `json:"dpr"`
		ADR                    float32 `json:"adr"`
		PctRoundsWithMK        float32 `json:"pct_rounds_with_mk"`
		OpeningKillsPerRound   float32 `json:"opening_kills_per_round"`
		WinPctAfterOpeningKill float32 `json:"win_pct_after_opening_kill"`
		Impact                 float32 `json:"impact"`
		KAST                   float32 `json:"kast"`
		Rating                 float32 `json:"rating"`
	}

	MatchCreate struct {
		FileHash               string            `json:"file_hash"`
		Map                    string            `json:"map"`
		TeamAName              string            `json:"team_a_name"`
		TeamBName              string            `json:"team_b_name"`
		TeamAScore             uint8             `json:"team_a_score"`
		TeamBScore             uint8             `json:"team_b_score"`
		TeamAScoreFirstHalf    uint8             `json:"team_a_score_first_half"`
		TeamBScoreFirstHalf    uint8             `json:"team_b_score_first_half"`
		TeamAScoreSecondHalf   uint8             `json:"team_a_score_second_half"`
		TeamBScoreSecondHalf   uint8             `json:"team_b_score_second_half"`
		TeamAOvertimeRoundsWon uint8             `json:"team_a_overtime_rounds_won"`
		TeamBOvertimeRoundsWon uint8             `json:"team_b_overtime_rounds_won"`
		PlayerData             map[string]Player `json:"player_data"`
	}

	MatchData struct {
		FileHash               string            `json:"file_hash"`
		Map                    string            `json:"map"`
		TeamAName              string            `json:"team_a_name"`
		TeamBName              string            `json:"team_b_name"`
		TeamAScore             uint8             `json:"team_a_score"`
		TeamBScore             uint8             `json:"team_b_score"`
		TeamAScoreFirstHalf    uint8             `json:"team_a_score_first_half"`
		TeamBScoreFirstHalf    uint8             `json:"team_b_score_first_half"`
		TeamAScoreSecondHalf   uint8             `json:"team_a_score_second_half"`
		TeamBScoreSecondHalf   uint8             `json:"team_b_score_second_half"`
		TeamAOvertimeRoundsWon uint8             `json:"team_a_overtime_rounds_won"`
		TeamBOvertimeRoundsWon uint8             `json:"team_b_overtime_rounds_won"`
		CreatedAt              string            `json:"created_at"`
		UpdatedAt              string            `json:"updated_at"`
		PlayerData             map[string]Player `json:"player_data"`
	}

	MatchDataOut struct {
		ID string `json:"id"`
	}
)
