// Package providing functionality to support individual note items in
// note.jlp
package note

type Note struct {
	ID        int
	NoteTitle string
	NoteBody  string // make this a byte slice?
}

func (n *Note) GenerateNote(id int, noteTitle, noteBody string) (*Note, error) {
	n.ID = id
	n.NoteTitle = noteTitle
	n.NoteBody = noteBody
	return n, nil
}

// Internal function to handle creation of a new note
func createNote() {

}

// Internal function to handle updates of an existing note
func updateNote() {

}

// Internal function to handle removing an existing note
func deleteNote() {

}
