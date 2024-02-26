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


request = Net::HTTP.get(uri)[1..-2]
result = request.tr("{}", "").split(",")
puts result


=begin
First attempt:

request = Net::HTTP.get(uri)
original_request = request
until request[0] == '{' do
    request = request[1..-1]
end

begin
    result = JSON.parse(request[0..-2])
rescue JSON::ParserError 
    # apparently, the error is caused by the "meanings" field.
    # "hello" is fine, but "beginning" is not.
    puts "Couldn't parse JSON request."
    puts "Original response: "
    puts original_request
else
    puts result
end
=end