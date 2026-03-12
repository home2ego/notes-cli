package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gitea.kood.tech/artemzhyrnyi/notes/constants"
)

func HandleSecurity(fileName string, r *bufio.Reader) bool {
	passFile := strings.Replace(fileName, ".txt", ".pass", 1)
	storedPass, err := os.ReadFile(passFile)

	if err != nil {
		fmt.Printf("\n%sNo password found. Set a new password for this file:%s\n", constants.Yellow, constants.Reset)
		fmt.Print("\n> ")

		newPass, _ := r.ReadString('\n')
		newPass = strings.TrimSpace(newPass)

		if len(newPass) > 0 {
			newPass += "\n"
		}

		_ = os.WriteFile(passFile, []byte(newPass), 0644)

		return true
	}

	fmt.Printf("\n%sEnter the password:%s\n", constants.Yellow, constants.Reset)
	fmt.Print("\n> ")
	attempt, _ := r.ReadString('\n')
	attempt = strings.TrimSpace(attempt)

	return attempt == strings.TrimSpace(string(storedPass))
}