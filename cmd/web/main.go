package main

// @title  Swagger API
// @version 0.0.1
// @schemes http

import (
	_ "shopline-question/docs"
	"shopline-question/initial"
	"shopline-question/router"
)

func main() {
	initial.Initial()
	r := router.NewRouter()
	err := r.Run(":8000")
	if err != nil {
		panic(err)
		return
	}
}
