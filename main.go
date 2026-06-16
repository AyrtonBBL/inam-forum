package main

import (
	"inam-forum/app"
)

func main() {
	
	application := app.InitApp()

	defer application.Db.Close()

	application.Run()
}