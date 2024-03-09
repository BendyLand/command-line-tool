package bland.user

import bland.notebook.*
import scala.io.StdIn

object User:
    var notebooks = List.empty[Notebook]

    def loop: Unit =
        println(
            "Please choose an action:\n" +
            "1.) Create new notebook\n2.) Write note\n3.) View notes\n4.) Delete note\n5.) Exit"
        )
        val input = StdIn.readLine().trim()
        input match
            case "1" => Notebook.createNotebook
            case "2" => Notebook.writeNewNote
            case "3" => Notebook.viewNotes
            case _ =>
                println("Shutting down notebook app...")

    def findNotebook: Option[Notebook] =
        println("Which notebook would you like to use?")
        showNotebooks
        try
            val notebookNum = StdIn.readInt().abs
            if notebookNum-1 >= notebooks.size then
                None
            else
                val nb = notebooks(notebookNum-1)
                Some(nb)
        catch
            case err: java.lang.NumberFormatException => 
                println("Invalid input")
                findNotebook

    def createNewNotebook(name: String) =
        notebooks = Notebook(name) +: notebooks

    def showNotebooks =
        for i <- 1 to notebooks.size do
            println(s"$i.) ${notebooks(i-1).name}")
        println()

    def init =
        greet
        loop

    def greet =
        println(
            "Welcome to the Scala Notebook App!\n" +
            "To get started, let's choose an action:\n" +
            "1.) Create\n2.) Exit"
        )
        var input = StdIn.readLine().trim()
        input match
            case "1" =>
                println("Please enter a name for your notebook:")
                val nbName = StdIn.readLine().trim()
                User.createNewNotebook(nbName)
                loop
            case _ =>
                println("Shutting down notebook app...")
