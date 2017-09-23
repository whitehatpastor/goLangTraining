package display

import (
	"fmt"
	//"bufio"
	//"os"
	"github.com/whitehatpastor/goLangTraining/rock_paper_scissors/inputs"
)



//test code
//func DrawWelcome(s int,m []string) {
//	fmt.Println("************************************************************")
//
//	for  i :=0; i<s; i++{
//		if i == s/2{
//
//		}
//		fmt.Println("*                                                          *")
//	}
//
//	fmt.Println("************************************************************")
//}

//Draws the welcome screen when the program first runs. Waits for a user input to continue to the next screen.
func DrawWelcome() {
	fmt.Println("************************************************************")
	fmt.Println("*                                                          *")
	fmt.Println("*                                                          *")
	fmt.Println("*              Welcome to Rock Paper Scissors              *")
	fmt.Println("*                                                          *")
	fmt.Println("*                                                          *")
	fmt.Println("************************************************************")
	inputs.WaitForInput("Press ENTER to Continue")
}


func DrawBattleground()  {
	fmt.Println("************************************************************")
	fmt.Println("*                                                          *")
	fmt.Println("*                                                          *")
	fmt.Println("*               Choose Wisley yout Weapon                  *")
	fmt.Println("*      Input the corresponding letter of your choice       *")
	fmt.Println("*        R = Rock     P = Paper    S = Scissors            *")
	fmt.Println("*                                                          *")
	fmt.Println("*                                                          *")
	fmt.Println("************************************************************")

}

func PlayAgain()  {
	fmt.Println("************************************************************")
	fmt.Println("*                                                          *")
	fmt.Println("*                                                          *")
	fmt.Println("*              Would you like to play again?               *")
	fmt.Println("*                          Y or N                          *")
	fmt.Println("*                                                          *")
	fmt.Println("*                                                          *")
	fmt.Println("************************************************************")

}
