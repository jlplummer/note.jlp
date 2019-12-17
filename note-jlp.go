package main

import (
	"fmt"

	"github.com/jlplummer/note.jlp/note"
)

func main() {
	var n = note.Note{}
	n.GenerateNote(0, "Note Title", "Note Body")
	fmt.Println("Your note:", "\n", n)

	var p Properties
	fp, p, err := p.OpenProperties()
	if err != nil {
		panic(err)
	}
	fmt.Println("Properties")
	fmt.Println("NotebookCntr:", p.NotebookCntr)
	fmt.Println("NoteCntr:", p.NoteCntr)
	fmt.Println("DataPath:", p.DataPath)

	for j := 0; j < 4; j++ {
		noteId, err := p.NewNoteId(fp)
		if err != nil {
			p.CloseProperties(fp)
			panic(err)
		}

		n = note.Note{noteId, "Note Title", "Note Body"}
		fmt.Println("Note:", n)
	}

	p.CloseProperties(fp)

}
