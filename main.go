package main

import (
	_ "bitcoin/config"
	"bitcoin/server"
)

func main() {
	var app server.App
	app.Run()
}
