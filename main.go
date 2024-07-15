package main

import (
	"go-todo/app/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Top)
	controllers.StartMainServer()
}
