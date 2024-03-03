package bland.notebook

import bland.note.Note

class Notebook:
    var notes = List.empty[Note]
    var currentId = 0

    def createNote(text: String): Note = 
        val note = Note(currentId, text)
        notes = note :: notes
        currentId += 1
        note

    def displayNotes: Unit = 
        println("Here are your notes: ")
        for i <- 1 to notes.size do
            println(s"$i.) ${notes(i-1).body}")