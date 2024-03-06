module Utils

open System

(*  
- 1/5 includes a mock utilization message (percentages; cpu, memory, disk space, etc.)
- 2/5 include a mock username
_*)

let greet = 
    printfn """
    Welcome to the Random Log Generator CLI!

    Please select which type of data you want:

        1.) Some jumbled data
        2.) A lot of organized data
    """

// let input = Console.ReadLine()
let generateData input =
    match input with
    | "1" -> printfn $"Generating some jumbled data..."
    | "2" -> printfn $"Generating a lot of organized data..."
    | _ -> printfn $"Invalid input"


// [timestamp] [client_ip] "GET /api/users HTTP/1.1" [status_code] [response_size]
// [timestamp] [client_ip] "POST /api/users HTTP/1.1" [status_code] [response_size]
// [timestamp] [client_ip] "PUT /api/users/123 HTTP/1.1" [status_code] [response_size]
// [timestamp] [client_ip] "DELETE /api/users/123 HTTP/1.1" [status_code] [response_size]