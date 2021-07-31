package main

import "myappEcho/app"

func main() {
	e := app.Init()

	e.Logger.Fatal(e.Start(":8080"))

}
