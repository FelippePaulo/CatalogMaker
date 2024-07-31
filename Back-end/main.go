package main

import (
	//"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"myproject/db"
)

type Catalogo struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

// type Catalog struct{
// 	Id int `json:"id"`
// 	Title string `json:"title"`
// 	Description string `json:"description"`
// 	Imglink string `json:"imgLink"`
// }

type Produto struct{
	Id string `json:"id"`
	IdCatalog string `json:"idCatalog"`
	Title string `json:"title"`
}

func getProducts(idCatalog string) []Produto{
	// Open the file
	file, err := os.Open("./products.json")
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		
	}
	defer file.Close()

	// Read the file's content
    byteValue, err := io.ReadAll(file)
    if err != nil {
        fmt.Printf("Failed to read file: %v\n", err)
       
    }

	// Parse the JSON data
    var responseP[] Produto 
    if err := json.Unmarshal(byteValue, &responseP); err != nil {
        fmt.Printf("Failed to parse JSON: %v\n", err)
        
    }
	
	var filteredProducts []Produto
	for _, produto := range responseP{
		if produto.IdCatalog == idCatalog{
			filteredProducts = append(filteredProducts, produto)
		}
	}
	return filteredProducts;
}

func main() {

	//DB connection
	err := db.Initialize()
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer db.DB.Close()

	//http request handler
	mux := http.NewServeMux()

	//post handler
	mux.HandleFunc(`/catalog`, func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Allow-Origin" , "*")
		//fmt.Printf(r.ParseForm().Error())
		fmt.Printf( "POST")
		fmt.Fprintf(w, "aaaaa")
		//fmt.Printf(obj)		
	})

	mux.HandleFunc("GET /catalog", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin" , "*")

		catalogs, err := db.GetCatalogs()
		if err != nil {
			http.Error(w, "Failed to retrieve catalogs", http.StatusInternalServerError)
			return
		}
		
		res, err := json.Marshal(catalogs)
		if err != nil{
			return
		}
		//w.Write([]byte(res))
		fmt.Fprintf(w, string(res))
	
	})

	mux.HandleFunc("GET /catalog/{id}", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin" , "*")
		id := r.PathValue("id")

		res, err := json.Marshal(getProducts(id))
		if err != nil{
			return
		}
		fmt.Fprintf(w, string(res))
	})

	//localhost server
	fmt.Println("ON")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}