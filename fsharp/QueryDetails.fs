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

module QueryDetails

open System

// "GET /api/users HTTP/1.1" 
// "POST /api/users HTTP/1.1" 
// "PUT /api/users/123 HTTP/1.1" 
// "DELETE /api/users/123 HTTP/1.1" 

(*  
[timestamp] represents the time at which the request was made.
[client_ip] represents the IP address of the client making the request.
The request line itself is logged in double quotes, including the HTTP method, the resource path, and the HTTP version.
[status_code] represents the HTTP status code returned by the server.
[response_size] represents the size of the response sent back to the client.
_*)

type RequestType = 
    | GET
    | POST
    | PUT
    | DELETE

let generateRandomIp() = 
    let rnd = Random()
    let nums = seq {for _ in 0..3 -> rnd.Next(256)} |> Seq.toList
    nums 
    |> List.map string
    |> String.concat "."

let chooseRandomRequestType() = 
    let options = [|GET; POST; PUT; DELETE|]
    let num = Random().Next(4)
    options[num]

