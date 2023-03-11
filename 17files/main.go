package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Work with files in GoLang")

	content := "This will go in a file - mukul98s"

	file, err := os.Create("./mukul98s.text")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./mukul98s.text")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text Data inside the file is \n", string(dataByte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
