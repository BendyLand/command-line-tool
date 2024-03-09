package bland.notebook

import scala.io.StdIn
import bland.note.*
import bland.user.*

object Notebook:
    def viewNotes =
        val notebook = User.findNotebook
        notebook match
            case Some(nb) => nb.displayNotes
            case None => println("Unable to find notebook")
        User.loop

    def writeNewNote = 
        val notebook = User.findNotebook
        notebook match
            case Some(nb) => nb.addNote
            case None => println("Unable to find notebook")
        User.loop

    def createNotebook = 
        println("Enter a name for your new notebook:")
        val name = StdIn.readLine()
        User.createNewNotebook(name)
        println(s"Successfully created notebook: '$name'")
        User.loop

    def addNote =
        println("Please select which notebook to write in:")
        User.showNotebooks
        try
            val notebookNum = StdIn.readInt()
            if notebookNum-1 >= User.notebooks.size then
                println("Couldn't find notebook")
            else
                println("Please enter your note: ")
                val nb = User.notebooks(notebookNum-1)
                val newNote = StdIn.readLine()
                nb.addNote
        catch
            case _: java.lang.NumberFormatException =>
                println("Invalid input")

class Notebook(val name: String):
    var notes = List.empty[Note]

    def addNote =
        println("Enter the text for your note:")
        val text = StdIn.readLine()
        val note = Note(text)
        notes = note :: notes

    def displayNotes =
        println("Here are your notes: ")
        for i <- 1 to notes.size do
            println(s"$i.) ${notes(i-1).body}")
        println()
