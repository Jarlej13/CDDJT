package main

import (
	"testing"
	"os"
)
//Testfunksjoner (TestFil, TestFil2, TestFil3) som sjekker om filene du oppgir er tilstede. Om de ikke er vil den returnere en error.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func TestFil(t *testing.T){
	r := Exists("template.html")
	if r == false {
		t.Error("File not found")
	}
}

func TestFil2(t *testing.T){
	r := Exists("template2.html")
	if r == false {
		t.Error("File not found")
	}
}

func TestFil3(t *testing.T){
	r := Exists("static/style.css")
	if r == false {
		t.Error("File not found")
	}
}
//Funksjon som tester om verdier for temperatur, nedbør og tilbakemelding fra kode samsvarer
func TestData(t *testing.T) {
	c, n, tilbakeMld := testData(5, 0, "oslo")
	if c > 0 {
		if tilbakeMld != "Det er kjølig ute. Ta på en jakke." {
			t.Error("Feil tilbakemelding")
		}

	}
	if n != 0 {
		t.Error("Feil nedbørsmengde")
	}

}
