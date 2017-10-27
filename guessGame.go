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

type text struct {
	Text string
	Guess string
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "guess.html")
}

func guessGameHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();
	guess := r.FormValue("guess")

	rand.Seed(int64(time.Now().Nanosecond()))

	if _, err := r.Cookie("target"); err != nil {
		var randNum = ((rand.Int() % 19) + 1)

		i, _ := strconv.Atoi(guess)


		if i == randNum{
			guess = "Correct!"
		}

		num := strconv.Itoa(randNum)


		cookie := http.Cookie{
			Name:  "target",
			Value: num,
		}

		http.SetCookie(w, &cookie)
	}

	m := text{Text: "Guess a number between 1 and 20: ", Guess: guess}
	t, _ := template.ParseFiles("guess.tmpl")
	t.Execute(w, m)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//http.HandleFunc("/", handler)
	http.HandleFunc("/guess", guessGameHandler)
	http.ListenAndServe(":8080", nil)
}