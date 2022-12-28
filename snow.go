package main

import (
	"fmt"

	"github.com/rpunt/dcc"
)


func main() {
	client := dcc.NewClient()
	client.BaseURL = "https://icanhazdadjoke.com"
	client.Headers["Accept"] = "application/json"
	response, err := client.Get("/")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(response.Body)
}
