# Note.jlp Design Document

Note.jlp is a learning project that I'm using to improve my Go knowledge along with playing with
web front end design.

## Project Structure
```bash
/ (Project Root)
    > config
        >> note-jlp.properties
    > static
        >> notes.css
        >> notes.html
    > data
    note-jlp.go
    note-jlp.md (This file)
    README.md
```

## API Structure

### Backend

##### Notebooks
| Endpoint | Purpose | Method |
| --- | --- | ---|
|[```/notebook/add```](#notebookadd)| Add Notebook | POST
|```/notebook/edit```| Edit Notebook | POST
|```/notebook/delete```| Delete Notebook| POST

##### Notes
|**Endpoint**|**Purpose**| Method |
|---|---|---|
|```/note/add```| Add Note to Notebook | POST
|```/note/edit```| Edit note in Notebook | POST
|```/note/delete```| Delete note from Notebook | POST

##### /notebook/add
