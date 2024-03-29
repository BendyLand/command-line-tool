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
    result = JSON.parse(request[0..-2]
rescue JSON::ParserError 
    # apparently, the error is caused by the "meanings" field.
    # "hello" is fine, but "beginning" is not.

    # UPDATE: JSON.parse appears to work when its input is the body of a 
    # get_response, rather than a get

    puts "Couldn't parse JSON request."
    puts "Original response: "
    puts original_request
else
    puts result
end
=end