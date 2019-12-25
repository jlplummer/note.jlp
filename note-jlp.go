package main

import (
	"fmt"

	"github.com/jlplummer/note.jlp/note"
)

func main() {
	var p Properties
	fp, p, err := p.OpenProperties()
	if err != nil {
		panic(err)
	}
	defer p.CloseProperties(fp)

	nbId, err := p.NewNotebookId(fp)
	if err != nil {
		panic(err)
	}

	var nb *note.Notebook = new(note.Notebook)
	nb, err = nb.GenerateNotebook(nbId, "Notebook Title", p.DataPath)
	if err != nil {
		panic(err)
	}

	var n *note.Note = new(note.Note)
	nId, err := p.NewNoteId(fp)
	if err != nil {
		panic(err)
	}

	fmt.Println("nb.NotebookPath", nb.NotebookPath)

	n, err = n.GenerateNote(nId, "Note Title", "Note Body", nb.NotebookPath)
	if err != nil {
		panic(err)
	}

	err = nb.AddNote(n)
	if err != nil {
		panic(err)
	}

	var n2 *note.Note = new(note.Note)
	nId2, err := p.NewNoteId(fp)
	if err != nil {
		panic(err)
	}

	n2, err = n2.GenerateNote(nId2, "Second Note Title", "2nd Note Body", nb.NotebookPath)
	if err != nil {
		panic(err)
	}

	err = nb.AddNote(n2)
	if err != nil {
		panic(err)
	}

	fmt.Println("nb after add", nb)

	n.NoteTitle = "Updated Title"
	success := nb.UpdateNote(n)
	if success != true {
		panic("cannot update note")
	}
	fmt.Println("nb after update", nb)

	fmt.Println("deleting", n2.ID)
	success = nb.DeleteNote(n2)
	if success != true {
		panic("cannot delete note")
	}

	fmt.Println("nb after delete", nb)
	/*
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

			// new() required to allocate memory for the pointer?
			var n *note.Note = new(note.Note)
			n, err = n.GenerateNote(noteId, "Note Title", "Note Body", p.DataPath)
			if err != nil {
				panic(err)
			}
			fmt.Println("Note:", n)
		}
	*/

}
