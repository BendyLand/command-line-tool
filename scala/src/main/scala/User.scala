package bland.user

import bland.notebook.*
import scala.io.StdIn

object User:
    var notebooks = List.empty[Notebook]

    def loop: Unit =
        println(
            "\nPlease choose an action:\n" +
            "\n1.) Create new notebook\n2.) Write note\n3.) View notes\n4.) Delete note\n5.) Delete notebook\n6.) Exit\n"
        )
        val input = StdIn.readLine().trim()
        input match
            case "1" => Notebook.createNotebook
            case "2" => Notebook.writeNewNote
            case "3" => Notebook.viewNotes
            case "4" => Notebook.deleteNote
            case "5" => User.deleteNotebook
            case ""  => User.loop
            case _ =>
                println("\nShutting down notebook app...")

    def findNotebook: Option[Notebook] =
        println("\nWhich notebook would you like to use?\n")
        displayNotebooks
        try
            val notebookNum = StdIn.readInt().abs
            if notebookNum-1 >= notebooks.size then
                None
            else
                val nb = notebooks(notebookNum-1)
                Some(nb)
        catch
            case _: java.lang.NumberFormatException => 
                println("\nInvalid input")
                findNotebook

    def deleteNotebook: Unit = 
        println("\nWhich notebook would you like to delete?\n")
        displayNotebooks
        val nbNum = StdIn.readInt().abs
        if nbNum-1 >= notebooks.size then
            println(s"\nInvalid selection. Please choose a number under ${notebooks.size}\n")
            deleteNotebook
        else
            val (firstHalf, secondHalf) = notebooks.splitAt(nbNum-1)
            notebooks = firstHalf ++ secondHalf.tail
            User.loop

    def createNewNotebook(name: String) =
        notebooks = Notebook(name) +: notebooks

    def displayNotebooks =
        for i <- 1 to notebooks.size do
            println(s"$i.) ${notebooks(i-1).name}")
        println()

    def init =
        greet
        loop

    def greet =
        println(
            "\nWelcome to the Scala Notebook App!\n" +
            "\nTo get started, let's choose an action:\n" +
            "\n1.) Create\n2.) Exit\n"
        )
        var input = StdIn.readLine().trim()
        input match
            case "1" =>
                println("\nPlease enter a name for your notebook:\n")
                val nbName = StdIn.readLine().trim()
                User.createNewNotebook(nbName)
                loop
            case _ =>
                println("\nShutting down notebook app...")
