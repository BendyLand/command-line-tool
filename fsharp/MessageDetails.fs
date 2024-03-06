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

let chooseRandomMessageType() = 
    let options = [|INFO; ERROR; WARN|]
    let num = Random().Next(3)
    options[num]

let chooseRandomMessageOrigin() = 
    let options = [|UserManagementService; Database; Security; SystemMonitor|]
    let num = Random().Next(4)
    options[num]
