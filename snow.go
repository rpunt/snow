package main

import (
	"fmt"

	"github.com/rpunt/simplehttp"
)


func main() {
	client := simplehttp.NewClient()
	client.BaseURL = "https://icanhazdadjoke.com"
	client.Headers["Accept"] = "application/json"
	response, err := client.Get("/")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(response.Body)
}
