package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductsResponse struct {
	Products []Product `json:"products"`
}

type ProductResponse struct {
	Product Product `json:"product"`
}

type WelcomeResponse struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/", productByIDHandler)

	fmt.Println("Starting server on port 8001")
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal(err)
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	response := WelcomeResponse{Message: "Welcome to CyanStack Commerce!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{ID: 1, Name: "Product 1"},
		{ID: 2, Name: "Product 2"},
	}

	response := ProductsResponse{Products: products}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func productByIDHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/products/")
	if path == "" {
		http.NotFound(w, r)
		return
	}

	productID, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := Product{ID: productID, Name: fmt.Sprintf("Product %d", productID)}
	response := ProductResponse{Product: product}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
