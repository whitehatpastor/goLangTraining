package main

import(
	//"fmt"
	"github.com/whitehatpastor/goLangTraining/rock_paper_scissors/display"
	"github.com/whitehatpastor/goLangTraining/rock_paper_scissors/inputs"

)

func winnerSelect(a string) int {
	//var c int
	if a == "r"{
		return 0
	} else if a == "p"{
		return 1
	}else {
		return 2
	}
}

func winTable(p int,c int) int {

}

func main() {
	var playerResponse string
	var cpuResponse string
	var p int
	var c int
	var winner int
	playGame := true

	display.DrawWelcome()

	for playGame {
		display.DrawBattleground()
		playerResponse = inputs.GameInput("Enter your Choice: ")
		p = winnerSelect(playerResponse)
		c = winnerSelect(cpuResponse)
		winner = winTable(p,c)
		display.PlayAgain()
		playGame = inputs.AskForConfirmation()
		}

}
