module MessageDetails

open Utils

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

let constructMessage () = 
    let messageType = chooseRandomMessageType()
    let message = 
        match messageType with
        | ERROR -> "Error message"
        | INFO -> "Info message"
        | WARN -> "Warning message"
    let origin = chooseRandomMessageOrigin()
    $"[%A{origin}] %s{message}"
