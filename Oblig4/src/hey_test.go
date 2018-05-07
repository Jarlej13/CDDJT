package main

import (
	"testing"
	"os"
)
//Testfunksjon som sjekker om filene du oppgir er tilstede. Om de ikke er vil den returnere en error.
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