module RequestDetails

open System
open Utils
open MyDateTime
open MessageDetails

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

let generateRandomIp () = 
    let nums = seq {for _ in 0..3 -> randomNumUnder 256} |> Seq.toList
    nums 
    |> List.map string
    |> String.concat "."

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

let generateRequestList count = 
    [|for _ in 0..count -> 
        let num = randomNumUnder 4
        match num with
        | 0 -> constructHttpRequest()
        | _ -> constructMessage ()
    |]

let generateRandomSqlQuery () = 
    let options = [|SELECT; INSERT; UPDATE; QueryType.DELETE|]
    let queryType = sample options
    match queryType with
    | SELECT -> "SELECT col1, col2 FROM table_name WHERE col2 IS NOT NULL"
    | INSERT -> "INSERT INTO table_name (col1, col2) VALUES (val1, val2)"
    | UPDATE -> "UPDATE table_name SET col1 = val1 WHERE col2 IS NOT NULL"
    | QueryType.DELETE -> "DELETE FROM table_name WHERE col2 IS NULL"


