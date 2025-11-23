package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv" 

	"github.com/eternalknight002/go-playground/internal/greeter"
)

type Response struct {
	Message string `json:"message"`
}

type IndexData struct {
	Title string
}


type DoubleResponse struct {
	Input  int `json:"input"`
	Result int `json:"result"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Lord"
	}

	msg := greeter.Greet(name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: msg})
}

//  Handler for the /double API
func doubleHandler(w http.ResponseWriter, r *http.Request) {
	inputStr := r.URL.Query().Get("number")
	input := 0

	if inputStr != "" {
		var err error
		// Try to parse the input string into an integer
		input, err = strconv.Atoi(inputStr)
		if err != nil {
			http.Error(w, "Invalid number format", http.StatusBadRequest)
			return
		}
	}
    // Call the new core logic function
	result := greeter.Double(input)

	w.Header().Set("Content-Type", "application/json")
    // Respond with the input and the calculated result
	json.NewEncoder(w).Encode(DoubleResponse{Input: input, Result: result})
}


func main() {
	// Serve static files under /static/
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Greet API
	http.HandleFunc("/greet", greetHandler)

	//  Double API
	http.HandleFunc("/double", doubleHandler)

	// Template rendering for root
	tmplPath := filepath.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := IndexData{
			Title: "Go Playground Server",
		}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "template error", http.StatusInternalServerError)
			log.Println("template execute error:", err)
		}
	})

	log.Println(" Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}