module Utils

open System
open System.IO
open QueryDetails
open MyDateTime

(*  
- 1/5 includes a mock utilization message (percentages; cpu, memory, disk space, etc.)
- 2/5 include a mock username
_*)

let greet() = 
    printfn """
    Welcome to the Random Log Generator CLI!

    Please select which type of data you want:

        1.) Some jumbled data
        2.) A lot of organized data
    """

let generateData input =
    match input with
    | "1" -> printfn $"Generating some jumbled data..."
    | "2" -> printfn $"Generating a lot of organized data..."
    | _ -> printfn $"Invalid input"

// [client_ip] [timestamp] "GET /api/users HTTP/1.1" [status_code] [response_size]
// [client_ip] [timestamp] "POST /api/users HTTP/1.1" [status_code] [response_size]
// [client_ip] [timestamp] "PUT /api/users/123 HTTP/1.1" [status_code] [response_size]
// [client_ip] [timestamp] "DELETE /api/users/123 HTTP/1.1" [status_code] [response_size]

// 127.0.0.1 - - [05/Mar/2024:12:30:45 +0000] "GET /index.html HTTP/1.1" 200 1234
// 2024-03-05T12:35:17 ERROR [UserManagementService] Error processing user registration request: User email already exists
// 2024-03-05T12:40:22 INFO [Database] Executing SQL query: SELECT * FROM users WHERE username='john_doe'
// 2024-03-05T12:45:55 WARN [Security] Suspicious login attempt detected: IP address 192.168.1.100, username 'admin'
// 2024-03-05T12:50:10 INFO [SystemMonitor] CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%

let constructHttpRequest() = 
    let requestType = chooseRandomRequestType()
    let path = "/index.html HTTP/1.1"
    let dateComponents = generateDateComponents()
    let timestamp = generateHttpStyleDate dateComponents
    let ip = generateRandomIp()
    let responseSize = Random().Next(800, 1601)
    $"%s{ip} %s{timestamp} \"%A{requestType} %s{path}\" 200 %d{responseSize}"

let generateRequestList count = 
    [|for _ in 0..count -> constructHttpRequest()|]

let writeRequestsToFile requestList = 
    File.AppendAllLines("sample_requests.txt", requestList)
