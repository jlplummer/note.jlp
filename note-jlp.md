# Note.jlp Design Document

Note.jlp is a learning project that I'm using to improve my Go knowledge along with playing with
web front end design.

## Project Structure
```bash
/ (Project Root)
    > config
        >> note-jlp.properties
    > data
    > static
        >> notes.css
        >> notes.html
    note-jlp.go
    note-jlp.md (This file)
    README.md
```

## Data Storage

**/data/notebooks/***{id}***/notes/***{noteId}*  
&nbsp;&nbsp;&nbsp;&nbsp;*id* - notebook identifier, obtained from [note-jlp.properties](#note-jlp.properties)  
&nbsp;&nbsp;&nbsp;&nbsp;*noteId* - individual note identifier, obtained from [note-jlp.properties](#note-jlp.properties)  

#### note-jlp.properties
````json
{
    "notebookCounter": 0,
    "noteCounter": 0
}
````

#### Sample Note
````json
{
    "id": int,
    "noteTitle": string,
    "noteBody": string
}
````

## API Structure

##### Notebooks
| Endpoint | Method  | Purpose |
| --- | --- | ---|
|[```/notebooks/{id}```](#notebooks)| GET | Retrieve Notebook data
|[```/notebooks/{id}```](#notebooks)| PUT | Add Notebook 
|[```/notebooks/{id}```](#notebooks)| POST | Update Notebook 

##### Notes
|Endpoint|Method| Purpose |
|---|---|---|
|[```/notebooks/{id}/notes/{noteId}```](#notes)| GET | Retrieve Note data
|[```/notebooks/{id}/notes/{noteId}```](#notes)| PUT | Add note in Notebook 
|[```/notebooks/{id}/notes/{noteId}```](#notes)| POST | Update note in Notebook 

##### /notebooks

##### /notes
