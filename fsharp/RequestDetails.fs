module RequestDetails

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
    let responseSize = randomNumBetween 800 1601
    $"%s{ip} - - %s{timestamp} \"%A{requestType} %s{path}\" 200 %d{responseSize}"

let chooseRandomQueryType () = 
    let options = [|SELECT; INSERT; UPDATE; QueryType.DELETE|]
    sample options
