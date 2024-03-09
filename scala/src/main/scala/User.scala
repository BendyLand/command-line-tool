package bland.user

import bland.notebook.*
import scala.io.StdIn

object User:
    var notebooks = List.empty[Notebook]

    def loop(): Unit = 
        println(
            "Please choose an action:\n" + 
            "1.) Create new notebook\n2.) Write note\n3.) View notes\n4.) Delete note\n5.) Exit"
        )
        val input = StdIn.readLine().trim()
        input match
            case "1" => 
                println("Enter a name for your new notebook:")
                val name = StdIn.readLine()
                createNotebook(name)
                println(s"Successfully created notebook: '$name'")
                loop()
            case "2" => 
                println("Please select which notebook to write in:")
                showNotebooks()
                try
                    val notebook = findNotebook()
                    notebook match
                        case Some(nb) => nb.writeNote()
                        case None => println("Unable to find notebook")
                catch 
                    case _: java.lang.NumberFormatException => 
                        println("Invalid input")
                loop()
            case "3" => 
                println("Please select which notebook to look through:")
                User.showNotebooks()
                try
                    val notebook = findNotebook()
                    notebook match
                        case Some(nb) => nb.displayNotes()
                        case None => println("Unable to find notes")
                catch 
                    case _: java.lang.NumberFormatException => 
                        println("Invalid input")
                loop()
            case _ => 
                println("Shutting down notebook app...")


    def findNotebook(): Option[Notebook] = 
        val notebookNum = StdIn.readInt().abs
        if notebookNum-1 >= User.notebooks.size then
            println("Couldn't find notebook")
            None
        else
            val nb = notebooks(notebookNum-1)
            Some(nb)

    def init() = 
        greet()
        loop()

    def createNotebook(name: String) = 
        notebooks = Notebook(name) +: notebooks

    def showNotebooks() = 
        for i <- 1 to notebooks.size do
            println(s"$i.) ${notebooks(i-1).name}")

    def greet() = 
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
                User.createNotebook(nbName)
                loop()
            case _ => 
                println("Shutting down notebook app...")