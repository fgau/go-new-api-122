package router

import (
	controllers "go-new-api-122/controller"
	"net/http"
)

func SetupRoutes() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("GET /users/{userID}", controllers.GetUsers)
	r.HandleFunc("GET /hello/{world}", controllers.HelloWorld)

	return r

}
