// Handles retrieving and updating the note.jlp properties file
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// where the note.jlp properties file is stored
const propertiesFile string = "config/note-jlp.properties"

// Structure that stores all property values from note-jlp.properties
type Properties struct {
	NotebookCntr int    `json:"notebookCounter"`
	NoteCntr     int    `json:"noteCounter"`
	DataPath     string `json:"dataPath"`
}

// Function handles opening and unmarshalling properties file into a struct.
//
// OpenProperties() does not close the opened file because other functions
// expect it to be open and will make modifications via a pointer.
func (p *Properties) OpenProperties() (*os.File, Properties, error) {
	propFile, err := os.Open(propertiesFile)
	if err != nil {
		return nil, *p, err
	}
	defer propFile.Close() //TODO: Figure out how to open this once for the duration of the program execution

	byteValue, err := ioutil.ReadAll(propFile)
	if err != nil {
		return nil, *p, err
	}

	json.Unmarshal(byteValue, &p)

	return propFile, *p, nil
}

// Closes the open Properties file
func (p *Properties) CloseProperties(fp *os.File) {
	fmt.Println("Closing", fp.Name())
	fp.Close()
}

func (p *Properties) NewNotebookId(fp *os.File) (int, error) {
	return 0, nil
}

func (p *Properties) NewNoteId(fp *os.File) (int, error) {
	p.NoteCntr++
	return p.NoteCntr, nil
}
