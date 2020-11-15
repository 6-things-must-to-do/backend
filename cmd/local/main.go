package main

import (
	"log"
	"github.com/6-things-must-to-do/server/internal"
	"github.com/gin-gonic/autotls"
 )

func main() {
	api := internal.GetAPI()
	log.Fatal(autotls.Run(api, "dev.sixthings.tech"))
}
