(*  
127.0.0.1 - - [05/Mar/2024:12:30:45 +0000] "GET /index.html HTTP/1.1" 200 1234
2024-03-05T12:35:17 ERROR [UserManagementService] Error processing user registration request: User email already exists
2024-03-05T12:40:22 INFO [Database] Executing SQL query: SELECT * FROM users WHERE username='john_doe'
2024-03-05T12:45:55 WARN [Security] Suspicious login attempt detected: IP address 192.168.1.100, username 'admin'
2024-03-05T12:50:10 INFO [SystemMonitor] CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%

- 2/5 include an IP address
- 1/5 includes an HTTP GET request
- 1/5 includes a SQL query
_*) 

module RequestDetails

open System
open Utils
open MyDateTime

type RequestType = 
    | GET
    | POST
    | PUT
    | DELETE

let generateRandomIp () = 
    let nums = seq {for _ in 0..3 -> randomNumUnder 256} |> Seq.toList
    nums 
    |> List.map string
    |> String.concat "."

let chooseRandomRequestType () = 
    let options = [|GET; POST; PUT; DELETE|]
    let num = randomNumUnder 4
    options[num]

let constructHttpRequest () = 
    let requestType = chooseRandomRequestType()
    let path = "/index.html HTTP/1.1"
    let dateComponents = generateDateComponents()
    let timestamp = generateHttpStyleDate dateComponents
    let ip = generateRandomIp()
    let responseSize = Random().Next(800, 1601)
    $"%s{ip} - - %s{timestamp} \"%A{requestType} %s{path}\" 200 %d{responseSize}"

let generateRequestList count = 
    [|for _ in 0..count -> constructHttpRequest()|]