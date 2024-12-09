package main

import (
    "fmt"
    "net/http"
    "log"
)

// helloHandler handles requests to the "/hello" URL path
func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Check if the URL path is "/hello"
    if r.URL.Path != "/hello" {
        // If not, return a 404 not found error
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
    // Check if the request method is "GET"
    if r.Method != "GET" {
        // If not, return a method not supported error
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }
    // Write "Hello!" to the response
    fmt.Fprintf(w, "Hello!")
}

// formHandler handles requests to the "/form" URL path
func formHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the form data from the request
    if err := r.ParseForm(); err != nil {
        // If there is an error, write the error to the response
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    // Indicate that the POST request was successful
    fmt.Fprintf(w, "POST request successful")
    
    // Retrieve the "name" and "address" form values
    name := r.FormValue("name")
    address := r.FormValue("address")

    // Write the form values to the response
    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
    // Create a file server to serve static files from the "./static" directory
    fileServer := http.FileServer(http.Dir("./static"))
    
    // Handle the root URL ("/") by serving static files using the file server
    http.Handle("/", fileServer)
    
    // Handle the "/form" URL by calling the formHandler function
    http.HandleFunc("/form", formHandler)
    
    // Handle the "/hello" URL by calling the helloHandler function
    http.HandleFunc("/hello", helloHandler)

		fmt.Printf("Server is running on port 8080\n")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
}