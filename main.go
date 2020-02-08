package main

import (
	"net/http"

	"github.com/SandeepMultani/restful_api_go/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
