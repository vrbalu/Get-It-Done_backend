package main

import (
	"GIT/controllers"
	"GIT/helpers"
)

func main() {
	r := controllers.SetupRouter()
	err := r.Run()
	if err != nil {
		helpers.Log.Error(err)
		return
	}
}
