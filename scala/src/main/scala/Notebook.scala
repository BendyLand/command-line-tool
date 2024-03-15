package bland.notebook

import scala.io.StdIn
import bland.note.*
import bland.user.*

object Notebook:
    def deleteNote = 
        val notebook = User.findNotebook
        notebook match
            case Some(nb) => nb.deleteNote
            case None     => println("\nUnable to find notebook")
        User.loop

    def viewNotes =
        val notebook = User.findNotebook
        notebook match
            case Some(nb) => nb.displayNotes
            case None     => println("\nUnable to find notebook")
        User.loop

    def writeNewNote = 
        val notebook = User.findNotebook
        notebook match
            case Some(nb) => nb.addNote
            case None     => println("\nUnable to find notebook")
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
                case Some(nb) => nb.addNote
                case None     => println("Couldn't find notebook")

class Notebook(val name: String):
    var notes = List.empty[Note]

    def deleteNote: Unit = 
        println("\nWhich note would you like to delete?")
        displayNotes
        try
            val noteNum = StdIn.readInt().abs
            if noteNum-1 >= notes.size then
                println(s"\nInvalid selection. Please choose a number under ${notes.size}.")
            else
                val (firstHalf, secondHalf) = notes.splitAt(noteNum-1)
                notes = firstHalf ++ secondHalf.tail
                User.loop
        catch
            case _: java.lang.NumberFormatException => 
                println("\nInvalid input")
                deleteNote

    def addNote =
        println("\nEnter the text for your note:\n")
        val text = StdIn.readLine()
        val note = Note(text)
        notes = note :: notes

    def displayNotes =
        println("\nHere are your notes:\n")
        for i <- 1 to notes.size do
            println(s"$i.) ${notes(i-1).body}")
        println()
