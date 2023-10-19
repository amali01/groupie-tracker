package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupietracker/apiprocess"
)

func RenderPage(w http.ResponseWriter, r *http.Request, data apiprocess.Data) {
	// Check if the request path is not the root path
	if r.URL.Path != "/" {
		HTTPErrorHandler(w, r, http.StatusNotFound)
		return
	}
	// Check if the request is not GET && NOT POST requests
	if r.Method != "GET" && r.Method != "POST" {
		HTTPErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodPost {
		for i := range data.Artists {
			data.Artists[i].DontDisplay = false
		}
		FilterCreation(w, r, data)
		FilterAlbum(w, r, data)
		Search(w, r, data)
		FilterConcerts(w, r, data)
		FilterNumbers(w, r, data)

	} else {
		for i := range data.Artists {
			data.Artists[i].DontDisplay = false
		}
	}

	files := []string{
		"views/html/index.html",
		"views/html/navigation.html",
		"views/html/artistscards.html",
		"views/html/footer.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		HTTPErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		HTTPErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func FilterCreation(w http.ResponseWriter, r *http.Request, data apiprocess.Data) {

	// Retrieve the value of the "textbox" form field
	cFrom, err1 := strconv.Atoi(r.FormValue("creationfrom"))
	cTo, err2 := strconv.Atoi(r.FormValue("creationto"))
	if cFrom == 0 || cTo == 0 {
		return
	}
	if err1 != nil {
		HTTPErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	if err2 != nil {
		HTTPErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	if cFrom > cTo {
		cTo, cFrom = cFrom, cTo
	}
	for i, artist := range data.Artists {
		if artist.CreationDate < cFrom || artist.CreationDate > cTo {
			data.Artists[i].DontDisplay = true
		}
	}
}

func FilterAlbum(w http.ResponseWriter, r *http.Request, data apiprocess.Data) {
	// Retrieve the value of the "textbox" form field
	cFrom, err1 := strconv.Atoi(r.FormValue("Albumfrom"))
	cTo, err2 := strconv.Atoi(r.FormValue("Albumto"))
	if cFrom == 0 || cTo == 0 {
		return
	}
	if err1 != nil {
		HTTPErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	
	if err2 != nil {
		HTTPErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if cFrom > cTo {
		cTo, cFrom = cFrom, cTo
	}

	for i := range data.Artists {
		firstAlbum, _ := strconv.Atoi(data.Artists[i].FirstAlbum[6:])
		if firstAlbum < cFrom || firstAlbum > cTo {
			data.Artists[i].DontDisplay = true
		}
	}

}

func Search(w http.ResponseWriter, r *http.Request, data apiprocess.Data) {
	Search := strings.ToLower(r.FormValue("query"))
	Search = strings.TrimSpace(Search)
	// fmt.Println("Change in Search value!", Search)

	if Search == "" {
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	for i := range data.Artists {
		SearchName := strings.ToLower(data.Artists[i].Name)
		SearchMember := strings.ToLower(strings.Join(data.Artists[i].Members, " "))
		SearchCreationDate := strconv.Itoa(data.Artists[i].CreationDate)
		SearchFirstAlbum := data.Artists[i].FirstAlbum

		Locations := []string{}

		for _, concert := range data.Artists[i].Conc {
			Locations = append(Locations, concert.Location)
		}
		SearchLocation := strings.Join(Locations, " ")

		if !strings.Contains(SearchName, Search) &&
			!strings.Contains(SearchMember, Search) &&
			!strings.Contains(SearchCreationDate, Search) &&
			!strings.Contains(SearchFirstAlbum, Search) &&
			!strings.Contains(SearchLocation, Search) {
			data.Artists[i].DontDisplay = true
		}

	}
}

func FilterConcerts(w http.ResponseWriter, r *http.Request, data apiprocess.Data) {
	filterLocation := strings.ToLower(r.FormValue("location"))

	if filterLocation == "" {
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	for i := range data.Artists {
		Locations := []string{}
		for _, concert := range data.Artists[i].Conc {
			Locations = append(Locations, concert.Location)
		}
		SearchLocation := strings.Join(Locations, " ")

		if !strings.Contains(SearchLocation, filterLocation) {
			data.Artists[i].DontDisplay = true
		}
	}
}

func FilterNumbers(w http.ResponseWriter, r *http.Request, data apiprocess.Data) {
	FilterNum := r.Form["NumMembers"]

	if len(FilterNum) == 0 {
		return
	}
	// fmt.Println("Change in FilterNumbers value!", FilterNum)

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	for i := range data.Artists {
		NumMembers := strconv.Itoa(len(data.Artists[i].Members))

		if !strings.Contains(strings.Join(FilterNum, " "), NumMembers) {
			data.Artists[i].DontDisplay = true
		}
	}
}
