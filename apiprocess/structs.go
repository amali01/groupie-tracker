package apiprocess

type Artist struct {
	ID           int      `json:"id"`
	Picture      string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Conc         Concerts
	DontDisplay  bool
}

type Concerts []struct {
	Location     string
	Country      string
	City         string
	Dates        []string
}

type List struct {
	Country string
	CitysList []string
}

type GeocodeResponse struct {
    Results []struct {
        Geometry struct {
            Lat float64 `json:"lat"`
            Lng float64 `json:"lng"`
        } `json:"geometry"`
    } `json:"results"`
}

type TmpAllConRel struct {
	Index []struct {
		Relation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type Data struct {
	Artists []Artist
	Lists []List
}

type TmpConcertsRel struct {
	Relation map[string][]string `json:"datesLocations"`
}
