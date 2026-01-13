package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)


type Note struct {
	Title string
	Content string
	Owner string
	CreatedAt time.Time

}

func New(title string, content string, owner string) *Note {
	return &Note{
		Title: title,
		Content: content,
		Owner: owner,
		CreatedAt: time.Now(),
	}

}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName = fileName + ".json"
	data_json, err := json.Marshal(note)
	if err != nil {
		return err 
	}
	return os.WriteFile(fileName, data_json, 0644)
}

func (note Note) DisplayNote() {
	fmt.Println(note.Title, note.Content,note.Owner)
}

