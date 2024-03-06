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


/// <summary>Returns a random integer up to the specified number (exclusive).</summary>
/// <param name="top">The exclusive upper bound for the number generation.</param>
let randomNumUnder top = 
    let rnd = Random()
    rnd.Next(top)

/// <summary>Returns a random integer between the specified lower and upper bounds.</summary>
/// <param name="bottom">The inclusive lower bound to be used in the random number generation range.</param>
/// <param name="top">The exclusive upper bound to be used in the random number generation range.</param>
let randomNumBetween bottom top = 
    let rnd = Random()
    rnd.Next(bottom, top)

/// <summary>Returns a random item from the provided collection.</summary>
/// <param name="coll">Collection from which to take a random element from.</param>
let sample coll = 
    let randomNum = randomNumUnder (Seq.length coll)
    Seq.item randomNum coll

    
