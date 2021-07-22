package main

import "fmt"

type user struct {
	Username string 
	FullName string
	ProfilePicUrl string
}

func main() {
	username := "Marino"
	fullName := "Kosaka Marino"
	profilePicUrl := "https://sakurazaka46.com/images/14/33f/b78b2083f09fb8208ec089504af14/400_320_102400.jpg"
	fmt.Println(username, fullName, profilePicUrl)

	u := user{
		Username: username,
		FullName: fullName,
		ProfilePicUrl: profilePicUrl,
	}
	// u.Username = username
	// u.FullName = fullName
	// u.ProfilePicUrl = profilePicUrl

	fmt.Printf("%#v\n",u)

	fmt.Println(u.Username)
	fmt.Println(u.FullName)
	fmt.Println(u.ProfilePicUrl)
}