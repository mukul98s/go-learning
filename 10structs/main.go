package main

import "fmt"

func main() {
  fmt.Println("Structs in GoLang")
  // inheritance is missing in GO;
  // No super for parent

  mukul := User{"Mukul Sharma", "mymukul@gmail.com", true, 22}
  fmt.Println(mukul)
  fmt.Printf("Mukul details are: %+v \n", mukul)
}

type User struct {
  Name   string
  Email  string
  Status bool
  Age    uint8
}
