package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Players struct {
	KiratBoli   OverStat `json:"Kirat_Boli"`
	NSNodhi     OverStat `json:"NS_Nodhi"`
	RRumhra     OverStat `json:"R_Rumhra"`
	ShashiHenra OverStat `json:"Shashi_Henra"`
}
type OverStat struct {
	Dot    int `json:"0"`
	Single int `json:"1"`
	Two    int `json:"2"`
	Three  int `json:"3"`
	Four   int `json:"4"`
	Five   int `json:"5"`
	Six    int `json:"6"`
	Out    int `json:"7"`
}

func main() {
	//keyArray := []string{"0", "1", "2", "3", "4", "5", "6", "out"}
	var players Players
	file, err := ioutil.ReadFile("input.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	err = json.Unmarshal(file, &players)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error : %+v ", err))
	}
	fmt.Println(fmt.Sprintf("Players info : %+v", players))
	kiratState := players.KiratBoli
	//nsNodhiState := players.NSNodhi
	//RumhraState := players.RRumhra
	//HenraState := players.ShashiHenra
	var kiratScore, nodhiScore, randScore, rumhraScore, henraScore int
	var kirat, nodhi, rumhra, henra bool
	kirat = true
	nodhi = true
	rumhra = false
	henra = false
	var kiratTurn, nodhiTurn, rumhraTurn, henraTurn bool
	kiratTurn = true
	nodhiTurn = true
	rumhraTurn = false
	henraTurn = false
	var kiratPower, nodhiPower, rumhraPower, henraPower int
	var kiratPlayedBalls, nodhiPlayedBalls, rumhraPlayedBalls, henraPlayedBalls int
	for j := 0; j < 4; j++ {
		for i := 0; i < 6; i++ {
			randScore = getRandom()
			if kirat && kiratTurn {
				kiratPower += getPlayerPercentage(randScore, kiratState)
				if randScore != 7 {
					kiratScore += randScore
				}
				if randScore == 7 {
					kirat = false
					if nodhi {
						rumhra = true
						rumhraTurn = true
					} else {
						henra = true
						henraTurn = true
					}
				} else if randScore%2 == 1 || randScore == 1 {
					if nodhi {
						nodhiTurn = true
					} else if rumhra {
						rumhraTurn = true
					} else if henra {
						henraTurn = true
					}
					kiratTurn = false
				} else {
					nodhiTurn = false
					kiratTurn = true
				}
			} else if nodhi && nodhiTurn {
				if randScore != 7 {
					nodhiScore += randScore
				}
				if randScore == 7 {
					nodhi = false
					if kirat {
						rumhra = true
						rumhraTurn = true
					} else {
						henra = true
						henraTurn = true
					}
				} else if randScore%2 == 1 || randScore == 1 {
					kiratTurn = true
					nodhiTurn = false
				} else if randScore != 7 {
					if kirat {
						kiratTurn = true
					} else if rumhra {
						rumhraTurn = true
					} else if henra {
						henraTurn = true
					}
					nodhiTurn = true
				}
			} else if rumhra && rumhraTurn {
				if randScore != 7 {
					rumhraScore += randScore
				}
				if randScore%2 == 1 || randScore == 1 {
					if kirat {
						kiratTurn = true
					} else if nodhi {
						nodhiTurn = true
					} else {
						henraTurn = true
					}
					rumhraTurn = false
				}
			} else if henra && henraTurn {
				if randScore != 7 {
					henraScore += randScore
				}
				if randScore%2 == 1 || randScore == 1 {
					if kirat {
						kiratTurn = true
					} else if nodhi {
						nodhiTurn = true
					} else {
						rumhraTurn = true
					}
					henraTurn = false
				}
			}
		}
	}
	fmt.Println("Kirat Score : ", kiratScore, "   Nodhi Score", nodhiScore, "  Rumhra Score ", rumhraScore, "  Henra Score ", henraScore)
}

func getRandom() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(8-0) + 0
}

func getPlayerPercentage(randScore int, playerState OverStat) int {
	var playerPercentage int
	switch randScore {
	case 0:
		{
			playerPercentage = playerState.Dot
		}
	case 1:
		{
			playerPercentage = playerState.Single
		}
	case 2:
		{
			playerPercentage = playerState.Two
		}
	case 3:
		{
			playerPercentage = playerState.Three
		}
	case 4:
		{
			playerPercentage = playerState.Four
		}
	case 5:
		{
			playerPercentage = playerState.Five
		}
	case 6:
		{
			playerPercentage = playerState.Six
		}
	case 7:
		{
			playerPercentage = playerState.Out
		}
	}
	return playerPercentage
}

func getMaxScoreofPlayer(playerState OverStat) int {
	var maxScore int
	if playerState.Five != 0 {
		maxScore += 5
	}
	if playerState.Four != 0 {
		maxScore += 4
	}
	if playerState.Single != 0 {
		maxScore++
	}
	if playerState.Six != 0 {
		maxScore += 6
	}
	if playerState.Three != 0 {
		maxScore += 3
	}
	if playerState.Two != 0 {
		maxScore += 2
	}
	return maxScore
}

func isthisWicket(players Players) {
	getMaxScoreofPlayer(players.KiratBoli)
}

func getPlayerRandomScore(playerState OverStat) {
	playerScoreBallMap := make(map[int]int)
	randScore := getRandom()
	var playerPower int
	if playerPower <= 60 {
		playerPower += getPlayerPercentage(randScore, playerState)
	} else {
		getPlayerPercentage(randScore, playerState)
	}
}

func getPlayerPredict(score, ballsCount int, playerStat OverStat) {
	if score == 60 {

	}
}
