package main

import (
	"fmt"
	"log"
	"net/http"

	"groupietracker/apiprocess"
	"groupietracker/controllers"
)

func main() {
	data, err := apiprocess.GetAllArtistsWithConcerts()
	if err != nil {
		fmt.Print("error")
		return
	}

	// Create an anonymous function that calls RenderPage with data   // Handler for the root path ("/") ..the home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.RenderPage(w, r, data)
	})

	http.HandleFunc("/css/", controllers.CSSHandler)        // Handler for serving CSS files
	http.HandleFunc("/images/", controllers.CSSHandler) // Handler for the "/GROUPIE.png" path

	http.HandleFunc("/artist/", func(w http.ResponseWriter, r *http.Request) {
		controllers.RenderArtistPage(w, r, data)
	})

	// Start the HTTP server
	fmt.Println("Server listening on port http://localhost:8080 ...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
