package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	http.HandleFunc("/", homeHandler)
	
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
