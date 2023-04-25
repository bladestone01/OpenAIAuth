package main

import (
	"fmt"
	"os"

	"github.com/acheong08/OpenAIAuth/auth"
)

func main() {
	auth := auth.NewAuthenticator(os.Getenv("OPENAI_EMAIL"), os.Getenv("OPENAI_PUID"), os.Getenv("PROXY"))
	err := auth.Begin()
	if err.Error != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		println("Embedded error: " + err.Error.Error())
		return
	}
	token, err := auth.GetAccessToken()
	if err.Error != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		println("Embedded error: " + err.Error.Error())
		return
	}
	fmt.Println(token)
}
