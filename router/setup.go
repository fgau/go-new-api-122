package router

import (
	controllers "go-new-api-122/controller"
	"net/http"
)

func SetupRoutes() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("GET /users/{userID}", controllers.GetUsers)
	r.HandleFunc("GET /hello/{world}", controllers.HelloWorld)
	r.HandleFunc("GET /favicon.ico", controllers.FaviconHandler)

	dir := http.Dir("./static")
	fs := http.FileServer(dir)
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	return r

}
