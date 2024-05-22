package routes

import (
	"net/http"
	"tugas-golang/handler"
	"tugas-golang/middleware"
)

func SetAPI() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /api/v1/get-user", middleware.MiddlewareBasicAuth(http.HandlerFunc(handler.GetUser)))
	mux.Handle("POST /api/v1/add-user", middleware.MiddlewareBasicAuth(http.HandlerFunc(handler.Storeuser)))

	mux.Handle("GET /api/v1/get-product", middleware.MiddlewareBasicAuth(http.HandlerFunc(handler.GetProducts)))
	mux.Handle("POST /api/v1/add-product", middleware.MiddlewareBasicAuth(http.HandlerFunc(handler.AddProduct)))

	return mux
}
