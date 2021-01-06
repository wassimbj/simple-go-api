package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/rs/cors"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products []Product

func getProducts(res http.ResponseWriter, req *http.Request) {
	reqMethod := req.Method
	if reqMethod != "" && strings.ToLower(reqMethod) != "get" {
		res.WriteHeader(400)
		res.Write([]byte("Invalid Method, expected GET"))
		return
	}
	res.Header().Set("Content-Type", "application/json")
	jsonP, _ := json.Marshal(products)
	res.Write(jsonP)
}

func createProduct(res http.ResponseWriter, req *http.Request) {
	reqMethod := req.Method
	if reqMethod == "" || strings.ToLower(reqMethod) != "post" {
		res.WriteHeader(400)
		res.Write([]byte("Invalid Method, expected POST"))
		return
	}

	body, _ := ioutil.ReadAll(req.Body)

	var data Product
	json.Unmarshal(body, &data)
	products = append(products, data)
	fmt.Println("Added product: ", products)
	res.WriteHeader(200)
	res.Write([]byte("Product created successfully"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", getProducts) // GET

	mux.HandleFunc("/create", createProduct) // POST

	// init server
	handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(":1234", handler)

	if err != nil {
		panic("Server is DOWN !!!")
	}

	fmt.Println("Server is running...")
}
