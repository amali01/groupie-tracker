package controllers

import (
	"html/template"
	"net/http"
	"strings"

	"groupietracker/apiprocess"
)

func RenderArtistPage(w http.ResponseWriter, r *http.Request, data apiprocess.Data) {
	// Check if the request is not GET requests
	if r.Method != "GET" {
		HTTPErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	artistName := strings.TrimPrefix(r.URL.Path, "/artist/")

	var artist apiprocess.Artist
	for _, a := range data.Artists {
		if a.Name == artistName {
			artist = a
			break
		}
	}
	
	if artist.ID == 0 {
		HTTPErrorHandler(w, r, http.StatusNotFound)
		return
	}

	artistPageData := struct {
		Artists []apiprocess.Artist
		Lists []apiprocess.List
		Artist apiprocess.Artist
	}{
		Artists: data.Artists,
		Lists: data.Lists,
		Artist: artist,
	}

	files := []string{
		"views/html/artist_page.html",
		"views/html/artist.html",
		"views/html/navigation.html",
		"views/html/footer.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		HTTPErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, artistPageData); err != nil {
		HTTPErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
