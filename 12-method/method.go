package main

import "fmt"

type User struct {
	Username string
	FullName string
	ProfilePicUrl string
}

func (u User) Info(){
	fmt.Printf("user name: %#v\n",u.Username)
	fmt.Printf("full name: %#v\n",u.FullName)
	fmt.Printf("profile picture url: %#v\n",u.ProfilePicUrl)
}

func main() {
	u := User{
		Username: "Marino",
		FullName: "Kosaka Marino",
		ProfilePicUrl: "https://sakurazaka46.com/images/14/33f/b78b2083f09fb8208ec089504af14/400_320_102400.jpg",
	}
	//info(u)
	u.Info()
}