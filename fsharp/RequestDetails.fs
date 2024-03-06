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

let chooseRandomMessageOrigin = 
    let options = [|SELECT; INSERT; UPDATE; QueryType.DELETE|]
    sample options
