=begin
Ruby project:
Dictionary lookup from:
    https://dictionaryapi.dev/

Example websites to use:
    https://example.com/
    https://jsonplaceholder.typicode.com/

Usage:
The basic syntax of a URL request to the API is shown below:

https://api.dictionaryapi.dev/api/v2/entries/en/<word>

As an example, to get definition of English word hello, you can send request to

https://api.dictionaryapi.dev/api/v2/entries/en/hello
=end

require "net/http"
require "json"

path_template = "https://api.dictionaryapi.dev/api/v2/entries/en/"

puts "Please enter a word to see its details: "
word = gets.chomp
uri = URI(path_template + word)


request = Net::HTTP.get_response(uri)
begin
    result = 
        request
            .body 
            .tr("{}", "") # remove the curly braces
            .split(",") # split at entries
            .join("\n")[1..-2] # format into a multiline string without surrounding brackets
    puts result
rescue 
    puts "Error parsing response"
end
