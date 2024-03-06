module MessageDetails

open System
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
    let num = randomNumUnder 3
    options[num]

let chooseRandomMessageOrigin () = 
    let options = [|UserManagementService; Database; Security; SystemMonitor|]
    let num = randomNumUnder 4
    options[num]

let constructMessage () = 
    let messageType = chooseRandomMessageType()
    let message = 
        match messageType with
        | ERROR -> "Error message"
        | INFO -> "Info message"
        | WARN -> "Warning message"
    let origin = chooseRandomMessageOrigin()
    $"[%A{origin}] %s{message}"
