package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
homeHandler reads the index.html file from disk, compiles it, and
sends it to the browser. When a user visits the home page, they 
receive the HTML form for entering text and selecting a banner.
*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Reject anything that isn't exactly "/"
	if r.URL.Path != "/" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}

	// Only allow GET
	if r.Method != http.MethodGet {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	// Parse and serve the template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
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
	// Parses the form data from the post request body
	// Must be called before reading any form values
	r.ParseForm()

	// Debugging lines to confirm the data that came in
	fmt.Println("In ascii art handler", r.Form)
	fmt.Println("Post form values:", r.PostForm)

	// Retrieve the form values
	inputText := r.Form.Get("inputText")
	banner := r.Form.Get("banner")

	result, err := AsciiArt(inputText, banner)
	// Gives the user a prope rreponse instead of a blank page
	// Pulls the error messages rom the AsciiArt function
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	// STOPPED HERE!!!!



}


func main() {
	http.HandleFunc("/", HomeHandler)
	
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
