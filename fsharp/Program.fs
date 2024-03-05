(* 
? Idea:
? Create a random Log generator to be used by a different program for data analysis.
? The "cli" part will be selecting the amount and type of data to include in the resulting file (see notes.txt)
_*)

open System

printfn """
Welcome to the Random Log Generator CLI!
Please select which type of data you're looking for: 
1.) Some jumbled data
2.) A lot of organized data 
"""

let input = Console.ReadLine()
let generateData input = 
    match input with
    | "1" -> printfn $"Generating some jumbled data..."
    | "2" -> printfn $"Generating a lot of organized data..."
    | _ -> printfn $"Invalid input"


generateData input
