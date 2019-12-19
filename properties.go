// Handles retrieving and updating the note.jlp properties file
package main

//TODO: Since you're just using WriteFile to update the properties, there is no need pass a file pointer around

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
func (p *Properties) OpenProperties() (*os.File, Properties, error) {
	propFile, err := os.OpenFile(propertiesFile, os.O_RDWR|os.O_TRUNC, 0744)
	if err != nil {
		return nil, *p, err
	}
	defer propFile.Close()

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
	fp.Sync()
	fp.Close()
}

// Increments NotebookCntr property, commits change to disk and returns the next counter
func (p *Properties) NewNotebookId(fp *os.File) (int, error) {
	p.NotebookCntr++
	p.writeProperties(fp)
	return p.NotebookCntr, nil
}

// Increments NoteCntr property, commits change to disk and returns the next counter
func (p *Properties) NewNoteId(fp *os.File) (int, error) {
	p.NoteCntr++
	p.writeProperties(fp)
	return p.NoteCntr, nil
}

func (p *Properties) writeProperties(fp *os.File) {
	b, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		fp.Close()
		panic(err)
	}

	fmt.Println("Writing", b)
	fmt.Println("fp", fp.Name())

	err = ioutil.WriteFile(propertiesFile, b, 0755)
	if err != nil {
		panic(err)
	}
}
