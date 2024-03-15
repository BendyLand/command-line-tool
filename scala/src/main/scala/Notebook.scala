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

    def addNote: Unit =
        val notebook = User.findNotebook
        val nb = 
            notebook match
                case Some(nb) => nb
                case None => 
                    println("Couldn't find notebook")
                    return
        println("\nPlease enter your note:\n")
        val newNote = StdIn.readLine()
        nb.addNote

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
