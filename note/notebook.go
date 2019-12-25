// Package providing functionality to support individual notebook items in
// note.jlp
package note

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Notebook struct {
	ID            int
	NotebookTitle string
	NotebookPath  string
	Notes         []Note
}

// Generates new notebook with the given ID by creating directory
// and creating .notebook file
func (n *Notebook) GenerateNotebook(id int, noteTitle, basePath string) (*Notebook, error) {
	n.ID = id
	n.NotebookTitle = noteTitle

	err := n.createNotebook(basePath)
	if err != nil {
		return nil, err
	}
	return n, nil
}

// Function to add an existing note to a notebook
func (n *Notebook) AddNote(note *Note) error {
	n.Notes = append(n.Notes, *note)

	b, err := json.MarshalIndent(n, "", " ")
	if err != nil {
		return err
	}

	fileToWrite := n.NotebookPath
	fileToWrite += "/"
	fileToWrite += strconv.Itoa(n.ID)
	fileToWrite += ".notebook"
	fmt.Println("fileToWrite", fileToWrite)
	err = ioutil.WriteFile(fileToWrite, b, 0755)
	if err != nil {
		return err
	}
	return nil
}

// Handles updating a specific note inside a notebook
// returns false if the note given cannot be found.
func (n *Notebook) UpdateNote(note *Note) bool {
	for j := 0; j < len(n.Notes); j++ {
		if note.ID == n.Notes[j].ID {
			n.Notes[j] = *note
			return true
		}
	}

	return false
}

// Handles deleting a specific note inside a notebook
// returns false if the note given cannot be found.
func (n *Notebook) DeleteNote(note *Note) bool {
	tmp := n.Notes[:0]
	success := false
	for j := 0; j < len(n.Notes); j++ {
		if note.ID == n.Notes[j].ID {
			success = true
			continue
		} else {
			tmp = append(tmp, n.Notes[j])
		}
	}
	n.Notes = tmp
	return success
}

// Internal function to handle creation of a new notebook
func (n *Notebook) createNotebook(basePath string) error {
	var dirToMake []string
	dirToMake = append(dirToMake, basePath)
	dirToMake = append(dirToMake, strconv.Itoa(n.ID))
	err := os.Mkdir(strings.Join(dirToMake, ""), 0755)

	n.NotebookPath = strings.Join(dirToMake, "")

	// may not be the best idea but tagging extra bits onto dirToMake and re-using the variable
	dirToMake = append(dirToMake, "/")
	dirToMake = append(dirToMake, strconv.Itoa(n.ID))
	dirToMake = append(dirToMake, ".notebook")

	b, err := json.MarshalIndent(n, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(strings.Join(dirToMake, ""), b, 0755)
	if err != nil {
		return err
	}
	return nil
}
