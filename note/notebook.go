// Package providing functionality to support individual notebook items in
// note.jlp
package note

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Notebook struct {
	ID            int
	NotebookTitle string
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
	b, err := json.MarshalIndent(n, "", " ")
	if err != nil {
		return err
	}

	var dirToMake []string
	dirToMake = append(dirToMake, basePath)
	dirToMake = append(dirToMake, strconv.Itoa(n.ID))
	err = os.Mkdir(strings.Join(dirToMake, ""), 0755)

	// may not be the best idea but tagging extra bits onto dirToMake and re-using the variable
	dirToMake = append(dirToMake, "/")
	dirToMake = append(dirToMake, strconv.Itoa(n.ID))

	err = ioutil.WriteFile(strings.Join(dirToMake, ""), b, 0755)
	if err != nil {
		return err
	}
	return nil
}
