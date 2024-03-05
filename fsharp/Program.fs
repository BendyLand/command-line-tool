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

// 127.0.0.1 - - [05/Mar/2024:12:30:45 +0000] "GET /index.html HTTP/1.1" 200 1234
// 2024-03-05T12:35:17 ERROR [UserManagementService] Error processing user registration request: User email already exists
// 2024-03-05T12:40:22 INFO [Database] Executing SQL query: SELECT * FROM users WHERE username='john_doe'
// 2024-03-05T12:45:55 WARN [Security] Suspicious login attempt detected: IP address 192.168.1.100, username 'admin'
// 2024-03-05T12:50:10 INFO [SystemMonitor] CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%

(*
Common themes:
√- All include a date. 4/5 are located at the front in the format YYYY-MM-DDThh:mm:ss. 
    - One is in square brackets using [DD/Mon/YYYY:hh:mm:ss +0000] (timezone offset).
- 4/5 include a message type, plus the origin of the log in square brackets
- 4/5 include a message explaining the log
- 2/5 include an IP address
- 1/5 includes an HTTP GET request
- 1/5 includes a SQL query
- 1/5 includes a mock utilization message (percentages; cpu, memory, disk space, etc.)
- 2/5 include a mock username
_*)

type MyDateTime = 
    {
        Year   : string
        Month  : string
        Day    : string
        Hour   : string
        Minute : string
        Second : string
    }
let createRecord (year::month::day::hour::minute::second::[]) =
    { Year = year; Month = month; Day = day; Hour = hour; Minute = minute; Second = second }

let generateDateComponents =
    let rnd = Random()
    let year = rnd.Next(1970, 2025)
    let month = rnd.Next(1, 13)
    let day = rnd.Next(1, 31)
    let hour = rnd.Next(1, 13)
    let minute = rnd.Next(1, 60)
    let second = rnd.Next(60)
    let elements = [year; month; day; hour; minute; second] |> List.map string
    let newElements = 
        elements
        |> List.map (fun x ->
            if String.length x < 2 then
                "0" + x
            else
                x
        )
    createRecord newElements

let generateDefaultDate (dateTime : MyDateTime) = 
    let {Year=year; Month=month; Day=day; Hour=hour; Minute=minute; Second=second} = dateTime
    $"%s{year}-%s{month}-%s{day}T%s{hour}:%s{minute}:%s{second}"

let components = generateDateComponents
let date = generateDefaultDate components
printfn $"%A{date}"
