package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv" 
    "strings" // Added for petnameHandler input processing

	"github.com/eternalknight002/go-playground/internal/greeter"
)

type Response struct {
	Message string `json:"message"`
}

type IndexData struct {
	Title string
}

type SubtractionResponse struct {
	N1     int `json:"n1"`
	N2     int `json:"n2"`
	Result int `json:"result"`
}


type DoubleResponse struct {
	Input  int `json:"input"`
	Result int `json:"result"`
}

// New struct for the Pet Name API response
type PetNameResponse struct {
	Adjective string `json:"adjective"`
	Animal    string `json:"animal"`
	PetName   string `json:"petName"`
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

// Handler for the /subtract API
func subtractHandler(w http.ResponseWriter, r *http.Request) {
	n1Str := r.URL.Query().Get("n1")
	n2Str := r.URL.Query().Get("n2")
	n1, n2 := 0, 0

	var err error
	// Parse n1, default to 0 if not present or invalid
	if n1Str != "" {
		n1, err = strconv.Atoi(n1Str)
		if err != nil {
			http.Error(w, "Invalid format for n1", http.StatusBadRequest)
			return
		}
	}

	// Parse n2, default to 0 if not present or invalid
	if n2Str != "" {
		n2, err = strconv.Atoi(n2Str)
		if err != nil {
			http.Error(w, "Invalid format for n2", http.StatusBadRequest)
			return
		}
	}
    
    // Core logic: n1 - n2
	result := greeter.Subtract(n1, n2)

	w.Header().Set("Content-Type", "application/json")
	// Respond with the two inputs and the calculated result
	json.NewEncoder(w).Encode(SubtractionResponse{N1: n1, N2: n2, Result: result})
}

// Handler for the /petname API (NEW FUN ENDPOINT)
func petNameHandler(w http.ResponseWriter, r *http.Request) {
	adjective := r.URL.Query().Get("adjective")
	animal := r.URL.Query().Get("animal")
	
    // Default values if parameters are missing
	if adjective == "" {
		adjective = "mysterious"
	}
	if animal == "" {
		animal = "shrimp"
	}
    
    // Ensure inputs are trimmed and lowercased before passing to logic
    adj := strings.TrimSpace(strings.ToLower(adjective))
    ani := strings.TrimSpace(strings.ToLower(animal))

	petName := greeter.GeneratePetName(adj, ani)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PetNameResponse{
		Adjective: adjective,
		Animal:    animal,
		PetName:   petName,
	})
}


func main() {
	// Serve static files under /static/
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Greet API
	http.HandleFunc("/greet", greetHandler)

	//  Double API
	http.HandleFunc("/double", doubleHandler)
    
    // Subtraction API
	http.HandleFunc("/subtract", subtractHandler)
    
    // New Pet Name API
	http.HandleFunc("/petname", petNameHandler)

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