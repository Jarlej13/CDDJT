package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	filename := os.Args[1]

	fi, err := os.Lstat(filename)
	//Her kommer en feilmelding dersom argumentet ikke er et filnavn i mappen.
	//Feilmeldingen sier "The system cannot find the file specified"
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Information about file <" + filename + ">\n")

	b := fi.Size()
	kb := b / 1024
	mb := kb / 1024
	gb := mb / 1024

	fmt.Printf("The filesize is: %d bytes, %d KB, %d MB, %d GB\n", b, kb, mb, gb)

	switch mode := fi.Mode(); {
	case mode.IsDir():
		fmt.Println("Is a directory")
	case !mode.IsDir():
		fmt.Println("Is not a directory")
	}

	switch mode2 := fi.Mode(); {
	case mode2.IsRegular():
		fmt.Println("Is a regular file")
	case !mode2.IsRegular():
		fmt.Println("Is not a regular file")
	}

	switch mode3 := fi.Mode(); {
	case mode3&os.ModePerm == 0777:
		fmt.Println("Has UNIX permission bits")
	case mode3&os.ModePerm != 0777:
		fmt.Println("Does not have UNIX permission bits")
	}

	switch mode4 := fi.Mode(); {
	case mode4&os.ModeAppend != 0:
		fmt.Println("Is append only")
	case mode4&os.ModeAppend == 0:
		fmt.Println("Is not append only")
	}

	switch mode5 := fi.Mode(); {
	case mode5&os.ModeDevice != 0:
		fmt.Println("Is a device file")
	case mode5&os.ModeDevice == 0:
		fmt.Println("Is not a device file")
	}

	switch mode8 := fi.Mode(); {
	case mode8&os.ModeSymlink != 0:
		fmt.Println("Is a symbolic link")
	case mode8&os.ModeSymlink == 0:
		fmt.Println("Is not a symbolic link")
	}
}
