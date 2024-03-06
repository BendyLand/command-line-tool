open System
open MyDateTime
open Utils
open MessageDetails
open QueryDetails

greet

let components = generateDateComponents
let date = generateDefaultDate components
printfn $"%A{date}"

let messageType = chooseRandomMessageType
printfn $"%A{messageType}"

let randomIp = generateRandomIp
printfn $"%s{randomIp}"

let messageOrigin = chooseRandomMessageOrigin
printfn $"%A{messageOrigin}"
