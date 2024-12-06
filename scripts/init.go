package api

import "fmt"

func Init() {
	fmt.Println("Initializing")
	web = Website{
		Template: "static/page.html",
	}
	GetAPI()
	CreateWebsite()

}
