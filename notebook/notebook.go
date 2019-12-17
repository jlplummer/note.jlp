// Package providing functionality to support individual notebooks in
// note.jlp
package notebook

import "jlplummer/note.jlp/note"

type Notebook struct {
	ID            int
	NotebookTitle string
	Notes         []note.Note
}
