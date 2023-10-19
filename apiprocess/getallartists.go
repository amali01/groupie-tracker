package apiprocess

import (
	"sort"
	"strings"
)

func GetAllArtistsWithConcerts() (Data, error) {
	tmprel := TmpAllConRel{}
	err := ParseAPI("https://groupietrackers.herokuapp.com/api/relation", &tmprel)
	if err != nil {
		return Data{}, err
	}

	// Creat variable of type Artists and ParseAPI unmarshals the response body into artists.
	var allArtists []Artist
	if err := ParseAPI("https://groupietrackers.herokuapp.com/api/artists", &allArtists); err != nil {
		return Data{}, err
	}

	var allLists []List
	// set concerts for all artists
	for i := range allArtists {
		// set concerts for each artist
		relMap := tmprel.Index[i].Relation
		// set the length of the struct by the number of the concerts
		artistConcerts := make(Concerts, len(relMap))
		conNum := 0
		for rawLoc, date := range relMap {
			artistConcerts[conNum].Location = strings.ToLower(rawLoc)

			// location := strings.ToLower(rawLoc)

			location := strings.Replace(strings.ToLower(rawLoc), "_", " ", -1)
			// artistConcerts[conNum].Location = location
			artistConcerts[conNum].Dates = date
			cityCountry := strings.Split(location, "-")
			if len(cityCountry) > 1 {
				City := strings.Title(cityCountry[0])
				artistConcerts[conNum].City = City
				Country := strings.ToUpper(cityCountry[1])
				artistConcerts[conNum].Country = Country
				// Add the location to the location list
				AddToList(City, Country, &allLists)
			}
			conNum++
		}
		allArtists[i].Conc = artistConcerts
	}
	SortList(&allLists)
	// To change the picture of Mamonas Assassinas Artist
	allArtists[20].Picture = "https://e-cdn-images.dzcdn.net/images/artist/94abb0f5039ec687e2f1413c96e64d68/264x264-000000-80-0-0.jpg"
	data := Data{
		Artists: allArtists,
		Lists:   allLists,
	}
	// fmt.Println(data)
	return data, nil
}

func AddToList(addCity string, addCountry string, allLists *[]List) {
	countryExist := false
	for i, co := range *allLists {
		if co.Country == addCountry {
			allCities := strings.Join(co.CitysList, " ")
			if !strings.Contains(allCities, addCity) {
				(*allLists)[i].CitysList = append((*allLists)[i].CitysList, addCity)
			}
			countryExist = true
		}
	}
	if !countryExist {
		list := List{
			Country:   addCountry,
			CitysList: []string{addCity},
		}
		*allLists = append(*allLists, list)
	}
}

func SortList(allLists *[]List) {
	sort.Slice(*allLists, func(i, j int) bool {
		return (*allLists)[i].Country < (*allLists)[j].Country
	})
	for i := range *allLists {
		sort.Strings((*allLists)[i].CitysList)
	}
}
