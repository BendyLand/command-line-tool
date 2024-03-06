(*  
127.0.0.1 - - [05/Mar/2024:12:30:45 +0000] "GET /index.html HTTP/1.1" 200 1234
2024-03-05T12:35:17 ERROR [UserManagementService] Error processing user registration request: User email already exists
2024-03-05T12:40:22 INFO [Database] Executing SQL query: SELECT * FROM users WHERE username='john_doe'
2024-03-05T12:45:55 WARN [Security] Suspicious login attempt detected: IP address 192.168.1.100, username 'admin'
2024-03-05T12:50:10 INFO [SystemMonitor] CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%

- 4/5 include a message type, plus the origin of the log in square brackets
- 4/5 include a message explaining the log
_*)

module MessageDetails

open System

type MessageType = 
    | ERROR
    | INFO
    | WARN

type MessageOrigin = 
    | UserManagementService
    | Database
    | Security
    | SystemMonitor

let chooseRandomMessageType = 
    let rnd = Random()
    let num = rnd.Next(3)
    match num with
        | 0 -> INFO
        | 1 -> ERROR
        | _ -> WARN

let chooseRandomMessageOrigin = 
    let rnd = Random()
    let num = rnd.Next(4)
    match num with
        | 0 -> UserManagementService
        | 1 -> Database
        | 2 -> Security
        | _ -> SystemMonitor