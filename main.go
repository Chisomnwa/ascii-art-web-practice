package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type TemplateData struct {
	Result string
}

/*
homeHandler reads the index.html file from disk, compiles it, and
sends it to the browser. When a user visits the home page, they 
receive the HTML form for entering text and selecting a banner.
*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET for this endpoint
	if r.Method != http.MethodGet {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	// Ensure the reuest is exactly for "/"
	if r.URL.Path != "/" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}

	// Parse and serve the template
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Println("Template parsing error:", err)
		http.Error(w, "500 internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template and check for errors
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Template excution error:", err)
		http.Error(w, "500 INternal SErver Error", http.StatusInternalServerError)
		return
	}
}

/*
AsciiArtHandler essentially processes the form submission and returns the result:
	- It recieves data from the POST request (the text the user typed and the
	  banner they selected),
	- Calls the AsciiArt() function to generate the ASCII art
	- Loads the index.html template
	- Inserts the genrated ASCII art into the template
	- Sends the complete HTML page (with the ASCII art displayed) back to the browser
*/
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Only POST requests are allowed on this endpoint
	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	// Ensure the request is exactly for "/ascii-art"
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	// Parses the form data from the post request body
	// Must be called before reading any form values
	err := r.ParseForm()
	if err != nil {
		// If parsing fails, the user sent malformed data
		fmt.Println("Form passing error:", err)
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	// Debugging lines to confirm the data that came in
	fmt.Println("In ascii art handler", r.Form)
	fmt.Println("Post form values:", r.PostForm)

	// Retrieve the form values
	inputText := r.Form.Get("inputText")
	banner := r.Form.Get("banner")

	result, err := AsciiArt(inputText, banner)
	// Pulls the error messages from the AsciiArt function and
	// Gives the user a proper reponse instead of a blank page
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tmp, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Tell the browser this is HTML content (good practice)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Execute the template and send it to the browser, and pass in the ASCII art result
	// Also, check if template execution fails
	err = tmp.Execute(w, &TemplateData{Result: result})
	if err != nil {
		fmt.Println("Template execution error:", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}


func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", AsciiArtHandler)
	
	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error startng server:", err)
	}
}





















// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		// w.Write([]byte("page not found"))
// 		http.NotFound(w, r)
// 		return
// 	}
// 	temp, _ := template.ParseFiles("./static/index.html")
// 	temp.Execute(w, nil)

// }

// func main() {

// 	http.HandleFunc("Get /", rootHandler)

// 	startRes := asciiArt("Our Server", "standard")
// 	fmt.Print(startRes)
// 	fmt.Println()

// 	fmt.Println("server starting on port 4000")
// 	err := http.ListenAndServe(":4000", nil)
// 	if err != nil {
// 		fmt.Println("error starting server:", err)

// 	}
// }
