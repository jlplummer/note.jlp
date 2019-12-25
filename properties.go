// Handles retrieving and updating the note.jlp properties file
package main

// TODO: Since you're just using WriteFile to update the properties, there is no need pass a file pointer around

import (
	"encoding/json"
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
func (p *Properties) OpenProperties() (Properties, error) {
	propFile, err := os.Open(propertiesFile)
	if err != nil {
		return *p, err
	}
	defer propFile.Close()

	byteValue, err := ioutil.ReadAll(propFile)
	if err != nil {
		return *p, err
	}

	json.Unmarshal(byteValue, &p)

	return *p, nil
}

// Increments NotebookCntr property, commits change to disk and returns the next counter
func (p *Properties) NewNotebookId() (int, error) {
	p.NotebookCntr++
	p.writeProperties()
	return p.NotebookCntr, nil
}

// Increments NoteCntr property, commits change to disk and returns the next counter
func (p *Properties) NewNoteId() (int, error) {
	p.NoteCntr++
	p.writeProperties()
	return p.NoteCntr, nil
}

func (p *Properties) writeProperties() {
	b, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(propertiesFile, b, 0755)
	if err != nil {
		panic(err)
	}
}
