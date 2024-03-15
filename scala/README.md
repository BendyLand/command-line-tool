# Scala Note Taking App

The purpose of this project is to help me practice writing command line tools of various forms. I chose Scala for this version of the project because it is my language of choice for many things, and should still be relatively forgiving for this size of project as I am learning. 

## Usage

1) Run the program with `sbt run` or start the sbt shell and type `run` separately.
2) You will be prompted with the option to create a notebook or exit. 
    - Exiting immediately basically just shuts down the app. It can be rerun from there.
3) After you name your notebook, you can interact by selecting one of the provided options.
    - While the app is running, you can:
        - Create new notebooks 
        - Write new notes in any of your notebooks 
        - View the notes in any of your notebooks
        - Delete any single notes from any of your notebooks 
        - Delete any of your notebooks (deleting all of the notes inside in the process) 
        - Exit the app
4) Once you exit the app, all of your previously created notes will be deleted. 
    - This app does NOT support memory between sessions (yet).
        - Since this is a practice project, it may never support memory, but I also might come back and add it once I am more knowledgeable about it. However, for the purpose of this project for the sake of practice, I am considering the functionality complete.