package main

import "fmt"

func main() {
  fmt.Println("Structs in GoLang")
  // inheritance is missing in GO;
  // No super for parent

  mukul := User{"Mukul Sharma", "mymukul@gmail.com", true, 22}
  fmt.Printf("Mukul details are: %+v \n", mukul)

	mukul.GetStatus()
	mukul.NewMail()
	
  fmt.Printf("Mukul details are: %+v \n", mukul)

}

type User struct {
  Name   string
  Email  string
  Status bool
  Age    uint8
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "random@random.com"
	fmt.Println("Email of this user is", u.Email)
}
