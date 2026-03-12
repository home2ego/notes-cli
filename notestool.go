package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gitea.kood.tech/artemzhyrnyi/notes/constants"
	"gitea.kood.tech/artemzhyrnyi/notes/utils"
)

func main() {
	// Validate arguments
	if len(os.Args) < 2 || strings.ToLower(os.Args[1]) == "help" {
		fmt.Print(constants.Clear)
		fmt.Printf("\nUsage: %s./todotool <filename>%s\n\n", constants.Purple, constants.Reset)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fileName := utils.FormatName(os.Args[1])

	fmt.Print(constants.Clear)

	// Validate password
	if !utils.HandleSecurity(fileName, reader) {
		fmt.Print(constants.Clear)
		fmt.Printf("\n%sAccess denied.%s\n\n", constants.Red, constants.Reset)
		return
	}

	fmt.Printf("\n%sWelcome to the notes tool!%s\n", constants.Purple, constants.Reset)

	for {
		fmt.Printf("\n%sSelect operation:%s\n", constants.Yellow, constants.Reset)
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")
		fmt.Print("\n> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			utils.ShowNotes(fileName)
		case "2":
			utils.AddNote(fileName, reader)
		case "3":
			utils.DeleteNote(fileName, reader)
		case "4":
			fmt.Print(constants.Clear)
			fmt.Printf("\n%sGoodbye!%s\n\n", constants.Purple, constants.Reset)
			os.Exit(0)
		case "":
			fmt.Print(constants.Clear)
			fmt.Printf("\n%sNot valid. Please provide an operation.%s\n", constants.Red, constants.Reset)
		default:
			fmt.Print(constants.Clear)
			fmt.Printf("\n%sNot valid. Please choose between options 1-4.%s\n", constants.Red, constants.Reset)
		}
	}
}

