module MyDateTime

open System

(*  
127.0.0.1 - - [05/Mar/2024:12:30:45 +0000] "GET /index.html HTTP/1.1" 200 1234
2024-03-05T12:35:17 ERROR [UserManagementService] Error processing user registration request: User email already exists
2024-03-05T12:40:22 INFO [Database] Executing SQL query: SELECT * FROM users WHERE username='john_doe'
2024-03-05T12:45:55 WARN [Security] Suspicious login attempt detected: IP address 192.168.1.100, username 'admin'
2024-03-05T12:50:10 INFO [SystemMonitor] CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%

√- All include a date. 4/5 are located at the front in the format YYYY-MM-DDThh:mm:ss. 
    - One is in square brackets using [DD/Mon/YYYY:hh:mm:ss +0000] (timezone offset).
_*)

type MyDateTime = 
    {
        Year   : string
        Month  : string
        Day    : string
        Hour   : string
        Minute : string
        Second : string
    }

let createRecord (year::month::day::hour::minute::second::[]) =
    { Year = year; Month = month; Day = day; Hour = hour; Minute = minute; Second = second }

let generateDateComponents() =
    let rnd = Random()
    let year = rnd.Next(1970, 2025)
    let month = rnd.Next(1, 13)
    let day = rnd.Next(1, 31)
    let hour = rnd.Next(1, 13)
    let minute = rnd.Next(1, 60)
    let second = rnd.Next(60)
    let elements = [year; month; day; hour; minute; second] |> List.map string
    let newElements = 
        elements
        |> List.map (fun x ->
            if String.length x < 2 then
                "0" + x
            else
                x
        )
    createRecord newElements

let convertMonthStyle monthNum =
    match monthNum with
        | "01" -> "Jan"
        | "02" -> "Feb"
        | "03" -> "Mar"
        | "04" -> "Apr"
        | "05" -> "May"
        | "06" -> "Jun"
        | "07" -> "Jul"
        | "08" -> "Aug"
        | "09" -> "Sep"
        | "10" -> "Oct"
        | "11" -> "Nov"
        | _ -> "Dec"

let generateDefaultDate (dateTime : MyDateTime) = 
    let {Year=year; Month=month; Day=day; Hour=hour; Minute=minute; Second=second} = dateTime
    $"%s{year}-%s{month}-%s{day}T%s{hour}:%s{minute}:%s{second}"

// [05/Mar/2024:12:30:45 +0000]
let generateHttpStyleDate (dateTime : MyDateTime)= 
    let {Year=year; Month=m; Day=day; Hour=hour; Minute=minute; Second=second} = dateTime
    let month = convertMonthStyle m
    $"[%s{day}/%s{month}/%s{year}:%s{hour}:%s{minute}:%s{second} +0000]"

