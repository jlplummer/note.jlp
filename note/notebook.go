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

func (n *Notebook) GenerateNotebook(id int, noteTitle, basePath string) (*Notebook, error) {
	n.ID = id
	n.NotebookTitle = noteTitle

	err := n.createNotebook(basePath)
	if err != nil {
		return nil, err
	}
	return n, nil
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

// Internal function to add an existing note to a notebook
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
