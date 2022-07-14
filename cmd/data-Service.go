package main

import (
	"fmt"
	"net/http"
	"tradeRevenueResultsService/controller"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func main() {
	savetodb := "mongodb"
	startService(savetodb)
}
func LogRequest(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func startService(db string) {
	router := chi.NewRouter()

	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(LogRequest)

	dataController := controller.NewDataController(db)
	dataController.WireRoutes(router)

	serverAddr := ":3335"

	for _, r := range router.Routes()[0].SubRoutes.Routes() {
		fmt.Printf(r.Pattern)
	}
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		fmt.Println("error")
	}
}
