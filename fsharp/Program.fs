open System
open MyDateTime
open Utils

let components = generateDateComponents
let date = generateDefaultDate components

greet
printfn $"%A{date}"
