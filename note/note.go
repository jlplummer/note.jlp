// Package providing functionality to support individual note items in
// note.jlp
package note

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
)

type Note struct {
	ID        int    `json:"id"`
	NoteTitle string `json:"noteTitle"`
	NoteBody  string `json:"noteBody"` // make this a byte slice?
}

func (n *Note) GenerateNote(id int, noteTitle, noteBody string, basePath string) (*Note, error) {
	n.ID = id
	n.NoteTitle = noteTitle
	n.NoteBody = noteBody

	err := n.createNote(basePath)
	if err != nil {
		return nil, err
	}
	return n, nil
}

// Internal function to handle creation of a new note
func (n *Note) createNote(basePath string) error {
	b, err := json.MarshalIndent(n, "", " ")
	if err != nil {
		return err
	}

	var fileToWrite []string
	fileToWrite = append(fileToWrite, basePath)
	fileToWrite = append(fileToWrite, strconv.Itoa(n.ID))
	fileToWrite1 := strings.Join(fileToWrite, "")

	err = ioutil.WriteFile(fileToWrite1, b, 0755)
	if err != nil {
		return err
	}
	return nil
}

// Internal function to handle updates of an existing note
func updateNote() {

}

// Internal function to handle removing an existing note
func deleteNote() {

}
