# goProblems2
Second set of Go Problems 

Curl

curl -i localhost:8080
The user enters this command in the command prompt (cmder), assuming the user has the nessessary files in the bin folder (Files curl.exe, curllib and curl-ca-bundle) which can be found on the curl website. Before this though the user much run the .go file by either converting it to an executionable file (.exe) or by typing "go run (File name)" in the command prompt in the directory where the file is. Once the curl command is executed then the console, the user should be able to view the contents by typing "http://localhost:8080" into the URL bar.

Bootstrap

The template used in these html files was found on the Bootstrap website at https://getbootstrap.com/docs/4.0/getting-started/introduction/#starter-template which has the nessesary stylesheet to style the html files. 

HTML/tmpl Files

This programme contains 2 html files, and a converted tmpl file, which are used to run the program on a webpage. The user first runs the guessGame.go file, then the user goes to the localhost webpage where it will display the index.html file. After clicking the new game link the user will be redirected to the guess.html page where the user will be asked to input a number between 1 and 20. Once the user guesses it will redirect them to the guess.tmpl page where it will display if they guessed correctly or not. If the user does guess correctly then a link will appear redirecting them to a new game.
