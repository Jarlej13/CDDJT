package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"html/template"
	"fmt"
	"runtime"
	"os/exec"
)
//Variabel som holder URL som vi kan endre med scanner
var URLen = ""
//Variabel som holder tilbakemelding ut fra nåværende temperatur
var tilbakeMld = ""
var input = ""

func main() {
	//Starter kode når man kobler til via localhost:1337/1
	openBrowser("http:\\localhost:1337/1")
	http.HandleFunc("/1", start)
	http.HandleFunc("/2", visData)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	if err := http.ListenAndServe(":1337", nil); err != nil {
		fmt.Print(err)
	}
}

//Lager struct basert på API hentet fra nett + Eget felt for tilbakemelding.
type Profil struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdated string  `json:"last_updated"`
		TempC       float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		WindKph    float64 `json:"wind_kph"`
		WindDir	   string     `json:"wind_dir"`
		PrecipMm   float64 `json:"precip_mm"`
		FeelslikeC float64 `json:"feelslike_c"`
	} `json:"current"`
	TilbakeMelding struct{
		TilbakeMld string `string:"tilbakeMld"`
	}
}
//Initialiserer structer
var profil Profil

//Funksjon for å endre data på felt i struct
func (f *Profil) setMld(melding string) {
	f.TilbakeMelding.TilbakeMld = melding
}

func returnStruct() Profil {
	return profil
}
//Når programmet starter får du opp nettside med inputfelt ved hjelp av denne koden. Den leser også inputen du gir.
func start(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template2.html")
		t.Execute(w, nil)
	}else {
		r.ParseForm()
	}
}
//Hoveddel av kode. Kobler til nettside, unmarshaler json data, parser det gjennom html fil som vises på nettleser
func visData(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("stedet")
	input = string(name)
	URLen = "http://api.apixu.com/v1/current.json?key=9c8d3cdcb75946a1bd272602182604&q=" + string(name)
	name = ""
	URL := string(URLen)
	res, err := http.Get(string(URL))
	if err != nil {
		log.Fatal(err)
	}


	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	jsonErr := json.Unmarshal(body, &profil)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	//Kjører kode for å hente temperatur data og tilbakemelding
	profil.hentData()
	//Setter tilbakemelding inn i struct via funksjon tidligere i koden
	profil.setMld(tilbakeMld)

	temp, err := template.ParseFiles("template.html")
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w, profil)
	if err != nil {
		log.Fatal(err)
	}

}

//Funksjon som henter og behandler data fra struct.
func (f *Profil) hentData() {
	c := f.Current.TempC
	n := f.Current.PrecipMm
	navn := f.Location.Name

	//Kjører if tester og lager responser ut fra det
	if c < 0 {
		if n > 1{
			tilbakeMld = "Det snør ute! (Med en sjanse for sludd)"
		}else{
			tilbakeMld = "Det er minusgrader ute. Ta på deg varmt tøy!"
		}
	} else if c < 10 {
		if n > 1{
			tilbakeMld = "Det regner og er kaldt ute. Kle deg godt."
		}else{
			tilbakeMld = "Det er kjølig ute. Ta på en jakke."
		}
	} else if c < 20 {
		if n > 1{
			tilbakeMld = "Det regner ute. Ta med en paraply."
		}else{
			tilbakeMld = "Det er rimelig varmt ute, men ikke gå for lettkledd."
		}
	} else if c < 30 {
		if n > 1{
			tilbakeMld = "Det regner ute. Ta med en paraply om du ikke skal på stranda."
		}else{
			tilbakeMld = "Det er varmt og godt ute! Ta på deg shorts og T-Skjorte og nyt været!"
		}
	} else {
		if n > 1{
			tilbakeMld = "Det regner ute, men det er så varmt det egentlig bare er deilig."
		}else{
			tilbakeMld = "Det er så varmt ute at du burde bade eller sitte inne bak airconditionen."
		}
	}
	if navn == "" {
		fmt.Println("Kjente ikke igjen sted. Vennligst start programmet på nytt med riktig stedsnavn. Du skrev inn " + input)
		tilbakeMld = "Kjente ikke igjen sted. Vennligst start programmet på nytt med riktig stedsnavn. Du skrev inn " + input
	}
}

//Åpner browser og navigerer til lenken oppgitt i main
func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}