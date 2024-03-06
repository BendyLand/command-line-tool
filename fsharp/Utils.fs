module Utils

open System
open System.IO

(*  
- 1/5 includes a mock utilization message (percentages; cpu, memory, disk space, etc.)
- 2/5 include a mock username
_*)

// 2024-03-05T12:35:17 ERROR [UserManagementService] Error processing user registration request: User email already exists
// 2024-03-05T12:40:22 INFO [Database] Executing SQL query: SELECT * FROM users WHERE username='john_doe'
// 2024-03-05T12:45:55 WARN [Security] Suspicious login attempt detected: IP address 192.168.1.100, username 'admin'
// 2024-03-05T12:50:10 INFO [SystemMonitor] CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%

let writeRequestsToFile requestList = 
    File.AppendAllLines("sample_requests.txt", requestList)

let randomNumUnder top = 
    let rnd = Random()
    rnd.Next(top)

let randomNumBetween bottom top = 
    let rnd = Random()
    rnd.Next(bottom, top)

