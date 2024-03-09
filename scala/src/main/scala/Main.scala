import bland.notebook.*

@main def run() = 
    val nb = Notebook()
    nb.createNote("This is a test note")
    nb.createNote("This is another test note")
    nb.createNote("This is a third test note")
    nb.displayNotes
    nb.deleteNote(2)
    nb.displayNotes


