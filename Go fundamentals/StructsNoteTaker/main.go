package main

import (
	"bufio"
	"fmt"
	"noteTaker/note"
	"os"
	"strings"
)


func main() {
	title, content := getNoteData()

	fmt.Println(title,content)

	var owner string = "joe"
	note := note.New(title,content,owner)
	
	note.DisplayNote()
	err:=note.Save()

	if err != nil{
		fmt.Println("saving to file gone wrong")
		return 
	}

	fmt.Println("Saving the note sucesss")

	
}
func getNoteData() (string, string) {
	
	title := getUserInput("Enter the title")
	content := getUserInput("Enter the content")
	return title, content
}

func getUserInput(prompt string) (userinput string) {
	fmt.Printf("%v ",prompt)

	reader:=bufio.NewReader(os.Stdin)
	userinput, err := reader.ReadString('\n')
	userinput = strings.TrimSuffix(userinput,"\n")
	userinput = strings.TrimSuffix(userinput,"\r")

	if err != nil {
		return ""
	}
	
	return userinput
}