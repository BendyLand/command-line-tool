module Utils

open System
open System.IO

/// <summary>Writes the contents of the provided collection to `sample_requests.txt`</summary>
/// <param name="requestList">The collection of elements to be written to the file</param>
let writeRequestsToFile requestList = 
    File.AppendAllLines("sample_requests.txt", requestList)

/// <summary>Returns a random integer up to the specified number (exclusive).</summary>
/// <param name="top">The exclusive upper bound for the number generation.</param>
let randomNumUnder top = 
    let rnd = Random()
    rnd.Next(top)

/// <summary>Generates a random number using a provided lower and upper bound.</summary>
/// <param name="bottom">The inclusive lower bound to be used in the random number generation range.</param>
/// <param name="top">The exclusive upper bound to be used in the random number generation range.</param>
/// <returns>A random integer between the specified lower and upper bounds</returns>
let randomNumBetween bottom top = 
    let rnd = Random()
    rnd.Next(bottom, top)

/// <summary>Selects one element from the provided collection at random.</summary>
/// <param name="coll">Collection to take the random element from.</param>
/// <returns>A random element from the provided collection.</returns>
let sample coll = 
    let randomNum = randomNumUnder (Seq.length coll)
    Seq.item randomNum coll
    
/// <summary>Creates a random IP address by generating five random numbers between 0 and 255 and joining them with '.'</summary>
/// <returns>A randomly generated IP address using IPv4 format.</returns>
let generateRandomIp () = 
    let nums = seq {for _ in 0..4 -> randomNumUnder 256} |> Seq.toList
    nums 
    |> List.map string
    |> String.concat "."

/// <summary>Generates full SQL query based on the input parameter</summary>
/// <param name="queryType">The type of SQL query to be returned</param>
/// <returns>A full SQL query of the same type as the input parameter, or a SELECT statement by default</returns>
let generateRandomSqlQuery queryType = 
    match queryType with
    | "INSERT" -> "INSERT INTO table_name (col1, col2) VALUES (val1, val2)"
    | "UPDATE" -> "UPDATE table_name SET col1 = val1 WHERE col2 IS NOT NULL"
    | "DELETE" -> "DELETE FROM table_name WHERE col2 IS NULL"
    | _ -> "SELECT col1, col2 FROM table_name WHERE col2 IS NOT NULL"
