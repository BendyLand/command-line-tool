(*  
2024-03-05T12:35:17 ERROR [UserManagementService] Error processing user registration request: User email already exists
2024-03-05T12:40:22 INFO [Database] Executing SQL query: SELECT * FROM users WHERE username='john_doe'
2024-03-05T12:45:55 WARN [Security] Suspicious login attempt detected: IP address 192.168.1.100, username 'admin'
2024-03-05T12:50:10 INFO [SystemMonitor] CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%
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

type QueryType = 
    | SELECT
    | INSERT
    | UPDATE
    | DELETE

let generateRandomIp () = 
    let nums = seq {for _ in 0..3 -> randomNumUnder 256} |> Seq.toList
    nums 
    |> List.map string
    |> String.concat "."

let chooseRandomRequestType () = 
    let options = [|GET; POST; PUT; RequestType.DELETE|]
    sample options

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

let generateRandomSqlQuery () = 
    let options = [|SELECT; INSERT; UPDATE; QueryType.DELETE|]
    let queryType = sample options
    match queryType with
    | SELECT -> "SELECT col1, col2 FROM table_name WHERE col2 IS NOT NULL"
    | INSERT -> "INSERT INTO table_name (col1, col2) VALUES (val1, val2)"
    | UPDATE -> "UPDATE table_name SET col1 = val1 WHERE col2 IS NOT NULL"
    | QueryType.DELETE -> "DELETE FROM table_name WHERE col2 IS NULL"


