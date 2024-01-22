package main

import (
	"adv/app"
	"adv/controller"
	"embed"
)

//go:embed resource
var resourses embed.FS

func main() {
	controller.Resource = resourses
	app.App1()
}