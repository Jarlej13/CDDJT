package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"html/template"
)

func main() {
	http.HandleFunc("/1", foo1)
	http.HandleFunc("/2", foo2)
	http.HandleFunc("/3", foo3)
	http.HandleFunc("/4", foo4)
	http.HandleFunc("/5", foo5)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

type Profil1 struct {
	Entries []struct {
		Kommune       string `json:"kommune"`
		Fylke         string `json:"fylke"`
		Navn          string `json:"navn"`
		SistEndret    string `json:"sistEndret"`
		Addresse      string `json:"addresse"`
		Aapningstider string `json:"aapningstider"`
		Lon           string `json:"lon"`
		ID            string `json:"id"`
		Lat           string `json:"lat"`
		URL           string `json:"url"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

type Profil2 struct {
	Entries []struct {
		Latitude  string `json:"latitude"`
		Navn      string `json:"navn"`
		Plast   string `json:"plast"`
		GlassMetall string `json:"glass_metall"`
		TekstilSko     string `json:"tekstil_sko"`
		Longitude string `json:"longitude"`
	} `json:"entries"`

}

type Profil3 struct {
	Entries []struct {
		CompanyCode int `json:"CompanyCode"`
		Dataset int `json:"Dataset"`
		StopCode string `json:"StopCode"`
		FullName string `json:"FullName"`
		Name string `json:"Name"`
		ShortName string `json:"ShortName"`
		Latitude float64 `json:"Latitude"`
		Longitude float64 `json:"Longitude"`
		Zone1 int `json:"Zone1"`
		Zone2 int `json:"Zone2"`
		TransferTime int `json:"TransferTime"`
		Transfer int `json:"Transfer"`
		BusStopType int `json:"BusStopType"`
		BusStopId int `json:"BusStopId"`
	} `json:"BusStops"`
}

type Profil4 struct {
	Entries[]struct {
		Kommune         string `json:"kommune"`
		Fylke			string `json:"fylke"`
		Navn	        string `json:"navn"`
	} `json:"entries"`
}

type Profil5 struct {
	Entries []struct {
		Name string`json:"name"`
		Code string `json:"code"`
	} `json:"countries"`

}

var profil1 Profil1
var profil2 Profil2
var profil3 Profil3
var profil4 Profil4
var profil5 Profil5

func foo1 (w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://hotell.difi.no/api/json/kmd/valglokaler/2015/forhand?")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	jsonErr := json.Unmarshal(body, &profil1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	temp, err := template.ParseFiles("templates/template1.html")
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w, profil1)
	if err != nil {
		log.Fatal(err)
	}

}

func foo2 (w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://hotell.difi.no/api/json/stavanger/miljostasjoner")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	jsonErr := json.Unmarshal(body, &profil2)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	temp, err := template.ParseFiles("templates/template2.html")
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w, profil2)
	if err != nil {
		log.Fatal(err)
	}

}

func foo3 (w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://sanntidsappservice-web-prod.azurewebsites.net/busstops?format=json")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	jsonErr := json.Unmarshal(body, &profil3)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	temp, err := template.ParseFiles("templates/template3.html")
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w, profil3)
	if err != nil {
		log.Fatal(err)
	}

}

func foo4 (w http.ResponseWriter, r *http.Request) {
	res, err := http.Get ("https://hotell.difi.no/api/json/difi/geo/kommune?fylke=14")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	jsonErr := json.Unmarshal(body, &profil4)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	temp, err := template.ParseFiles("templates/template4.html")
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w,profil4)
	if err != nil {
		log.Fatal(err)
	}

}

func foo5 (w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://api.nobelprize.org/v1/country.json")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	jsonErr := json.Unmarshal(body, &profil5)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	temp, err := template.ParseFiles("templates/template5.html")
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w, profil5)
	if err != nil {
		log.Fatal(err)
	}

}
