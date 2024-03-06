module Startup

open System
open Utils
open RequestDetails
open MessageDetails

let generateRequestList count = 
    [|for _ in 0..count -> 
        let num = randomNumUnder 4
        match num with
        | 0 -> constructHttpRequest ()
        | _ -> constructMessage ()
    |]

let rec init () = 
    printfn """
    Welcome to the Random Log Generator CLI!

    Please enter the amount of random log files to generate: 
    """
    let rec loop () = 
        let input = Console.ReadLine() 
        match Int32.TryParse(input) with
        | true, num -> 
            printfn "Generating data..."
            let requests = generateRequestList num
            writeRequestsToFile requests
            printfn "Data written to file `sample_requests.txt`"
        | _-> 
            printfn "Invalid input. Please enter a number."
            loop()
    loop()
