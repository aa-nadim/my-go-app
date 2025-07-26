package routers

import (
	"my-go-app/controllers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", controllers.Main)
}
