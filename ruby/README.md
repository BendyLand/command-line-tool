# Ruby Dictionary Lookup

The purpose of this project was to give me a more intermediate-level project to work on; in this case, one that takes the form of a command-line tool. I chose to implement it in Ruby to start since it is a fairly forgiving language which would allow me to quickly and easily prototype with different techniques. 

## Usage

To use the CLI, run `ruby script.rb` and follow the prompts printed in the console:
1) First, you will enter a word to look up.
    - Due to the way data is fetched from [the API](https://dictionaryapi.dev/), if you spell the word incorrectly or enter an otherwise unrecognized word, the program will print an error message and exit here. 
2) Next, you will select which piece of information you would like to see. 
    - Invalid inputs here may result in unexpected results. You should still get a response, but it may not be the information you are expecting. 
    - Overflowing numbers will wrap around the options, but all other inputs appear to result in `sourceUrls`.
3) That's it! Your chosen information about the word you entered will be written to a file called `word_data.txt` in this same directory. 

### Additional Considerations

I wasn't quite sure where to end this project. Parsing the JSON data further is possible, depending on which information you chose. I just didn't see much reason to do so, since all I would be doing is displaying it anyway. I figured by writing it to a file instead, at least the data wouldn't disappear right away, even if it's slightly more cryptic looking. 

I might come back and do something more meaningful with the data at a later time. 

