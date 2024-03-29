=begin
Ruby project:
Dictionary lookup from:
    https://dictionaryapi.dev/

Usage:
The basic syntax of a URL request to the API is shown below:

https://api.dictionaryapi.dev/api/v2/entries/en/<word>

As an example, to get definition of English word hello, you can send request to

https://api.dictionaryapi.dev/api/v2/entries/en/hello
=end

require "net/http"
require "json"

class Dictionary
    def choose_data_to_view(data)
        if data.nil? 
            return 
        end
        fields = ["word", "phonetics", "meanings", "license", "sourceUrls"]
        puts "Please select which field you would like to know about:"
        for i in 0...fields.length do 
            puts "#{i+1}.) #{fields[i]}"
        end
        begin 
            choice = gets.chomp.to_i % 5 - 1
            puts "Writing #{fields[choice]} for '#{data['word']}' to file: 'word_data.txt'"
            data[fields[choice]]
        rescue
            puts "Invalid option"
        end
    end

    def enter_word()
        puts "To get started, please enter a word: "
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
        end
    end
end

dict = Dictionary.new

puts "Welcome to my Dictionary Lookup Tool, where you can get relevant information about a word of your choice!"
word = dict.enter_word()
data = dict.get_word_data(word)
parsed_data = dict.parse_word_data(data)
result = dict.choose_data_to_view(parsed_data)

result_string = JSON.pretty_generate(result)

unless result_string == "null"
    File.open("word_data.txt", "w") do |file| 
        file.write("Retrieved data for '#{word}':\n")
        file.write(result_string) 
    end
else
    puts "Failed to retrieve data for '#{word}'. File will not change."
end

