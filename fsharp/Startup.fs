module Startup

open System
open Utils
open RequestDetails

let rec init () = 
    printfn """
    Welcome to the Random Log Generator CLI!

    Please select which type of data you want:

        1.) Some jumbled data
        2.) A lot of organized data
    """
    let rec loop () = 
        let input = Console.ReadLine()
        match input with
        | "1" -> 
            printfn "Generating some jumbled data..."
            printfn "(This doesn't do anything yet)"
        | "2" ->
            printfn "Generating a lot of organized data..."
            let requests = generateRequestList 10000
            writeRequestsToFile requests
            printfn "Data written to file `sample_requests.txt`"
        | _ -> 
            printfn "Invalid input. Please enter '1' or '2'"
            loop()
    loop()