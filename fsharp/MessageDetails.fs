module MessageDetails

open Utils
open MyDateTime

type MessageType = 
    | ERROR
    | INFO
    | WARN

type MessageOrigin = 
    | UserManagementService
    | Database
    | Security
    | SystemMonitor

let chooseRandomMessageType () = 
    let options = [|INFO; ERROR; WARN|]
    sample options

let chooseRandomMessageOrigin () = 
    let options = [|UserManagementService; Database; Security; SystemMonitor|]
    sample options

let generateErrorMessage messageOrigin = 
    match messageOrigin with
    | UserManagementService -> "Error processing user registration request: User email already exists"
    | Database -> "Error executing SQL Query"
    | Security -> "Permission Denied: You do not have access this content"
    | SystemMonitor -> "Not enough disk space to perform task"

let generateInfoMessage messageOrigin =
    match messageOrigin with
    | UserManagementService -> "Successfully created user account!"
    | Database -> 
        let origin = chooseRandomMessageOrigin().ToString()
        let query = generateRandomSqlQuery origin
        $"Executing SQL query: %s{query}"
    | Security -> 
        let ip = generateRandomIp()
        $"Successful login detected: IP address %s{ip}, username 'admin'"
    | SystemMonitor -> "CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%"

// let generateWarnMessage origin =


(*  
2024-03-05T12:35:17 ERROR [UserManagementService] Error processing user registration request: User email already exists
2024-03-05T12:40:22 INFO [Database] Executing SQL query: SELECT * FROM users WHERE username='john_doe'
2024-03-05T12:45:55 WARN [Security] Suspicious login attempt detected: IP address 192.168.1.100, username 'admin'
2024-03-05T12:50:10 INFO [SystemMonitor] CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%
_*) 
let constructMessage () = 
    let messageType = chooseRandomMessageType()
    let dateComponents = generateDateComponents()
    let timestamp = generateDefaultStyleDate dateComponents
    let origin = chooseRandomMessageOrigin()
    let message = 
        match messageType with
        | ERROR -> generateErrorMessage origin
        | INFO -> generateInfoMessage origin
        | WARN -> "Warning message"
    $"%s{timestamp} %A{messageType} [%A{origin}] %s{message}"
