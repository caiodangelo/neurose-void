package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func createOrder(ctx context.Context, w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Creating order")
}

func updateOrder(ctx context.Context, w http.ResponseWriter, r *http.Request){
	orderId := pat.Param(ctx, "id")
	fmt.Fprintf(w, "Updating order %s", orderId)
}

func readOrder(ctx context.Context, w http.ResponseWriter, r *http.Request){
	orderId := pat.Param(ctx, "id")
	fmt.Fprintf(w, "Reading order %s", orderId)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/order/:id"), readOrder)
	mux.HandleFuncC(pat.Post("/order"), createOrder)
	mux.HandleFuncC(pat.Put("/order/:id"), updateOrder)

	http.ListenAndServe("localhost:8000", mux)
}