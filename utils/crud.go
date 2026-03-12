package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"gitea.kood.tech/artemzhyrnyi/notes/constants"
)

// Load notes
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

// Save notes
func saveNotes(fileName string, notes []string) {
	content := strings.Join(notes, "\n")

	if len(content) > 0 {
		content += "\n"
	}

	_ = os.WriteFile(fileName, []byte(content), 0644)
}

// Show notes
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

// Add a new note
func AddNote(fileName string, r *bufio.Reader) {
	fmt.Printf("\n%sEnter the note text:%s\n", constants.Yellow, constants.Reset)
	fmt.Print("\n> ")
	text, _ := r.ReadString('\n')
	text = strings.TrimSpace(text)

	if text == "" {
		fmt.Print(constants.Clear)
		fmt.Printf("\n%sCannot add an empty note.%s\n", constants.Red, constants.Reset)
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04")
	entry := fmt.Sprintf("%s [%s]", text, timestamp)

	notes := loadNotes(fileName)
	notes = append(notes, entry)
	saveNotes(fileName, notes)
	fmt.Printf("\n%sNote added.%s\n", constants.Green, constants.Reset)
}

// Delete a note
func DeleteNote(fileName string, r *bufio.Reader) {
	notes := loadNotes(fileName)

	if len(notes) == 0 {
		fmt.Printf("\n%s%q is empty.%s\n", constants.Yellow, fileName, constants.Reset)
		return
	}

	fmt.Println("\nEnter the number of note to remove or 0 to cancel:")
	fmt.Print("\n> ")

	number, _ := r.ReadString('\n')
	number = strings.TrimSpace(number)
	idx, err := strconv.Atoi(number)

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