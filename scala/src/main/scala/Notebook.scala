package bland.notebook

import scala.io.StdIn
import bland.note.*
import bland.user.*

object Notebook:
    def viewNotes =
        val notebook = User.findNotebook
        notebook match
            case Some(nb) => nb.displayNotes
            case None => println("\nUnable to find notebook")
        User.loop

    def writeNewNote = 
        val notebook = User.findNotebook
        notebook match
            case Some(nb) => nb.addNote
            case None => println("\nUnable to find notebook")
        User.loop

    def createNotebook = 
        println("\nEnter a name for your new notebook:\n")
        val name = StdIn.readLine()
        User.createNewNotebook(name)
        println(s"\nSuccessfully created notebook: '$name'")
        User.loop

    def addNote =
        println("\nPlease select which notebook to write in:\n")
        User.displayNotebooks
        try
            val notebookNum = StdIn.readInt()
            if notebookNum-1 >= User.notebooks.size then
                println("\nCouldn't find notebook")
            else
                println("\nPlease enter your note:\n")
                val nb = User.notebooks(notebookNum-1)
                val newNote = StdIn.readLine()
                nb.addNote
        catch
            case _: java.lang.NumberFormatException =>
                println("\nInvalid input")

class Notebook(val name: String):
    var notes = List.empty[Note]

    def addNote =
        println("\nEnter the text for your note:\n")
        val text = StdIn.readLine()
        val note = Note(text)
        notes = note :: notes

    def displayNotes =
        println("\nHere are your notes:\n")
        for i <- 1 to notes.size do
            println(s"$i.) ${notes(i-1).body}")
