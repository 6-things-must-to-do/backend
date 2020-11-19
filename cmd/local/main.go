package main

import (
	"github.com/6-things-must-to-do/server/internal"
	"github.com/gin-gonic/autotls"
	"log"
)

func main() {
	api := internal.GetAPI()
	log.Fatal(autotls.Run(api, "dev.sixthings.tech"))
}
