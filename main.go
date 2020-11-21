package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	//The Load method reads the .env file and loads the set variables into the
	//environment so that they can be accessed through the os.Getenv() method
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()
	fmt.Println("server running on port 3000!!!")

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
