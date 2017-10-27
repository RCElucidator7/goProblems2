//Ryan Conway
//Adapted from https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.4.html
package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//struct that gathers the values to display on the page
type text struct {
	Text string
	Guess string
	NewGame string
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "guess.html")
}

func guessGameHandler(w http.ResponseWriter, r *http.Request) {
	//Parse the input that the user inputs in the webpage
	r.ParseForm();
	guess := r.FormValue("guess")

	//Generates a random number using the time import
	rand.Seed(int64(time.Now().Nanosecond()))

	//Generates a random number between 1 and 20
	var number = ((rand.Int() % 19) + 1)

	//Sets the cookie with a name of target
	cookie, err := r.Cookie("target")

	//If theres no errors it converts the string value to an int
	if err == nil {
		number, _= strconv.Atoi(cookie.Value)
	}

	//Initalise newgame for when the user guesses correctly
	newgame := ""

	//Assigns the parsed input into a variable and converts it from string to int
	userInput, _ := strconv.Atoi(guess)

	//if statements to check if the user guesses to high, low or correctly
	if userInput == number{
		guess = "Thats Correct!"
		newgame = "Click here to start a new game!"
	}else if userInput < number{
		guess = "The number " + strconv.Itoa(userInput) + " was too low! Try another number!"
	}else{
		guess = "The number " + strconv.Itoa(userInput) + " was too high! Try another number!"
	}

	//Assigns the cookie values to a variable
	cookie = &http.Cookie{
		Name:  "target",
		Value: strconv.Itoa(number),
	}

	//Sets the cookie on the webpage
	http.SetCookie(w, cookie)

	//Puts the data into the guess.tmpl file
	m := text{Text: "Please Guess a number between 1 and 20: ", Guess: guess, NewGame: newgame}
	t, _ := template.ParseFiles("guess.tmpl")
	t.Execute(w, m)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/guess", guessGameHandler)
	http.ListenAndServe(":8080", nil)
}