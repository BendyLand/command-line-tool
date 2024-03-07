module MessageDetails

open Utils
open MyDateTime
open RequestDetails

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
    | Database -> "There was a problem executing the SQL Query."
    | Security -> "Permission Denied: You do not have access to this content."
    | SystemMonitor -> "Not enough disk space to perform task."

let generateInfoMessage messageOrigin =
    match messageOrigin with
    | UserManagementService -> "Successfully created user account!"
    | Database -> 
        let origin = chooseRandomQueryType().ToString()
        printfn $"{origin}"
        let query = generateRandomSqlQuery origin
        $"Executing SQL query: %s{query}"
    | Security -> 
        let ip = generateRandomIp()
        $"Successful login: IP address %s{ip}, username 'admin'"
    | SystemMonitor -> "CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%"

let generateWarningMessage origin =
    match origin with
    | UserManagementService -> "Password may not contain enough variety. Proceed with caution."
    | Database -> "SQL query may be missing rows. Please double check your input data."
    | Security -> 
        let ip = generateRandomIp()
        $"Suspicious login attempt detected: IP address %s{ip}, username 'admin'"
    | SystemMonitor -> "Disk space utilization at 80%. Consider adding more storage or removing unneeded files."

let constructMessage () = 
    let messageType = chooseRandomMessageType()
    let dateComponents = generateDateComponents()
    let timestamp = generateDefaultStyleDate dateComponents
    let origin = chooseRandomMessageOrigin()
    let message = 
        match messageType with
        | ERROR -> generateErrorMessage origin
        | INFO -> generateInfoMessage origin
        | WARN -> generateWarningMessage origin
    $"%s{timestamp} %A{messageType} [%A{origin}] %s{message}"
