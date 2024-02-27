# This file will be for all the things I try throughout this project
# that do not make it into the final version.  

=begin
First attempt:

request = Net::HTTP.get(uri)
original_request = request
until request[0] == '{' do
    request = request[1..-1]
end

begin
    result = JSON.parse(request[0..-2]u
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