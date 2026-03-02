package utils

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"gitea.kood.tech/artemzhyrnyi/notes/constants"
)

func loadNotes(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		
		if line != "" {
			notes = append(notes, line)
		}
	}

	return notes
}

func saveNotes(fileName string, notes []string) {
	content := strings.Join(notes, "\n")

	if len(content) > 0 {
		content += "\n"
	}

	_ = os.WriteFile(fileName, []byte(content), 0644)
}

func ShowNotes(fileName string) {
	notes := loadNotes(fileName)

	if len(notes) == 0 {
		fmt.Printf("\n%sNo notes found in %q.%s\n", constants.Yellow, fileName, constants.Reset)
		return
	}

	fmt.Printf("\n%sNotes in %q:%s\n", constants.Yellow, fileName, constants.Reset)

	for i, note := range notes {
		fmt.Printf("%s%03d%s - %s\n", constants.Cyan, i+1, constants.Reset, note)
	}
}

func AddNote(fileName string, r *bufio.Reader) {
	fmt.Printf("\n%sEnter the note text:%s\n", constants.Yellow, constants.Reset)
	fmt.Print("\n> ")

	input, _ := r.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Print(constants.Clear)
		fmt.Printf("\n%sCannot add an empty note.%s\n", constants.Red, constants.Reset)
		return
	}
	
	notes := loadNotes(fileName)
	notes = append(notes, input)
	saveNotes(fileName, notes)
	fmt.Printf("\n%sNote added.%s\n", constants.Green, constants.Reset)
}

func DeleteNote(fileName string, r *bufio.Reader) {
	notes := loadNotes(fileName)

	if len(notes) == 0 {
		fmt.Printf("\n%s%q is empty.%s\n", constants.Yellow, fileName, constants.Reset)
		return
	}

	fmt.Println("\nEnter the number of note to remove or 0 to cancel:")
	fmt.Print("\n> ")

	input, _ := r.ReadString('\n')
	input = strings.TrimSpace(input)
	idx, err := strconv.Atoi(input)

	if err != nil || idx < 0 || idx > len(notes) {
		fmt.Print(constants.Clear)
		fmt.Printf("\n%sInvalid ID. No changes made.%s\n", constants.Red, constants.Reset)
		return
	}

	if idx == 0 {
		fmt.Printf("\n%sDeletion cancelled.%s\n", constants.Yellow, constants.Reset)
		return
	}

	notes = append(notes[:idx-1], notes[idx:]...)
	saveNotes(fileName, notes)
	fmt.Printf("\n%sNote %03d removed.%s\n", constants.Green, idx, constants.Reset)
}