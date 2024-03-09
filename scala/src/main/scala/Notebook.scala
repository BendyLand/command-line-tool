package bland.notebook

import scala.io.StdIn
import bland.note.*
import bland.user.*

class Notebook(val name: String):
    var notes = List.empty[Note]
    var currentId = 0

    def writeNote() =
        println("Enter the text for your note:")
        val text = StdIn.readLine()
        val note = Note(currentId, text)
        notes = note :: notes
        currentId += 1

    def viewNotes(nb: Notebook) = 
        nb.displayNotes()

    def addNote() = 
        println("Please select which notebook to write in:")
        User.showNotebooks()
        try
            val notebookNum = StdIn.readInt()
            if notebookNum-1 >= User.notebooks.size then
                println("Couldn't find notebook")
            else
                println("Please enter your note: ")
                val nb = User.notebooks(notebookNum-1)
                val newNote = StdIn.readLine()
                nb.writeNote()
        catch 
            case _: java.lang.NumberFormatException => 
                println("Invalid input")

    def displayNotes() =
        println("Here are your notes: ")
        for i <- 1 to notes.size do
            println(s"$i.) ${notes(i-1).body}")

    def deleteNote(id: Int): Boolean =
        var newList = List.empty[Note]
        notes.foreach { note =>
            if note.id != id-1 then
                newList = note :: newList
        }
        var result = true
        if newList.size == notes.size then
            println("Warning, possible invalid note ID. Please double check your notebook.")
            result = false
        notes = newList
        result