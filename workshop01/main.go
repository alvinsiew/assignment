package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Message struct {
   Quote string
   Repo string
}

func main() {
   port := os.Getenv("PORT")
   rand.Seed(time.Now().Unix())
   quotes := []string{
    "Logic will get you from A to B. Imagination will take you everywhere.",
    "There are 10 kinds of people. Those who know binary and those who don't.",
    "There are two ways of constructing a software design. One way is to make it so simple that there are obviously no deficiencies and the other is to make it so complicated that there are no obvious deficiencies.",
    "It's not that I'm so smart, it's just that I stay with problems longer.",
    "It is pitch dark. You are likely to be eaten by a grue.",
   }
   // n := rand.Int() % len(quotes)

   templates := template.Must(template.ParseFiles("templates/index.html"))

   http.Handle("/static/",
      http.StripPrefix("/static/",
         http.FileServer(http.Dir("static"))))

   http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {
      n := rand.Int() % len(quotes)
      message := Message{quotes[n], "Repository:"}

      if err := templates.ExecuteTemplate(w, "index.html", message); err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
      }
   })

   fmt.Println("Listening");
   fmt.Println(http.ListenAndServe(":"+port, nil));
}