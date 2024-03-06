open MyDateTime
open Utils
open MessageDetails
open QueryDetails

greet()

let requests = generateRequestList 100
writeRequestsToFile requests