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

class Dictionary
    def enter_word()
        puts "Please enter a word: "
        word = gets.chomp
        unless word.nil?
            word
        else
            ""
        end
    end

    def get_word_data(word)
        path_template = "https://api.dictionaryapi.dev/api/v2/entries/en/"
        uri = URI(path_template + word)
        request = Net::HTTP.get_response(uri)
        if request.message == "OK"
            request.body
        else
            "Error: Could not retrieve data for word: '#{word}'"
        end
    end

    def parse_word_data(data)
        begin 
            result = JSON.parse(data)
            result[0]
        rescue
            puts "Error: Could not parse word data"
            {}
        end
    end
end

dict = Dictionary.new
word = dict.enter_word()
data = dict.get_word_data(word)
result = dict.parse_word_data(data)

puts result