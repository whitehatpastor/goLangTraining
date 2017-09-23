package inputs

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strings"

)

// askForConfirmation uses Scanln to parse user input. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user. Typically, you should use fmt to print out a question
// before calling askForConfirmation. E.g. fmt.Println("WARNING: Are you sure? (yes/no)")
func AskForConfirmation() bool {
	fmt.Print("Play Again? : ")
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	response = strings.ToLower(response)
	if response == "y" {
		return true
	} else if response == "n" {
		return false
	}else {
		fmt.Println(`Please type "y" for yes or "n" for no and then press enter:`)
		return AskForConfirmation()
	}
}


//Displays a message passed to it and waits for the user to respond.
func WaitForInput(s string)  {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print(s)
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(sentence))
	}
}

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
//func posString(slice []string, element string) int {
//	for index, elem := range slice {
//		if elem == element {
//			return index
//		}
//	}
//	return -1
//}

// containsString returns true if slice contains element
//func containsString(slice []string, element string) bool {
//	return !(posString(slice, element) == -1)
//}


func GameInput(s string) string{
	var response string
	fmt.Println(s)
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	response = strings.ToLower(response)
	return response

}