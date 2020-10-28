package main

import "github.com/6-things-must-to-do/server/internal"

func main() {
	api := internal.GetAPI()
	api.Run(":4000")
}