package main

import "myappEcho/routes"

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))

}
