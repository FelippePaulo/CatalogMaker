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
	mux.HandleFunc(`POST /catalog`, func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Allow-Origin" , "*")
		
		fmt.Printf("Request Method: %s\n", r.Method)
		// Ensure this is a POST request
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        var data map[string]interface{}
        decoder := json.NewDecoder(r.Body)
		
        if err := decoder.Decode(&data); err != nil {
            http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
            return
        }
		
		var catalog db.Catalog
		
        // Acesse os valores do JSON
        catalog.Title = data["item1"].(string)
        catalog.Description = data["item2"].(string)
		catalog.Imglink = data["item3"].(string)


		db.AddCatalog(catalog)

        fmt.Fprintf(w, "Received JSON data")	
	})

	// change catalog handler
	mux.HandleFunc(`/catalog/{id}`, func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == http.MethodOptions {
			// Handle preflight request
			w.WriteHeader(http.StatusOK)
			return
		}
	})

	// put req catalog
	mux.HandleFunc(`PUT /catalog/{id}`, func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodPut {
			var data map[string]interface{}
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&data); err != nil {
				http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
				return
			}
			
			var catalog db.Catalog 
			
			id := r.PathValue("id")

			catalog.Id = id
			catalog.Title = data["item1"].(string)
			catalog.Description = data["item2"].(string)
			catalog.Imglink = data["item3"].(string)

			db.AlterCatalog(catalog)
		}

	})

	// delete catalog
	mux.HandleFunc("DELETE /catalog/{id}", func(w http.ResponseWriter, r *http.Request){
		 // Handle CORS
		 w.Header().Set("Access-Control-Allow-Origin", "*")
		 w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
		 w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	 
		 // Handle DELETE request
		 if r.Method == http.MethodDelete {
			fmt.Println("DELETE request")
			id := r.PathValue("id")
			db.DeleteCatalog(id)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Item deleted successfully"))
		 } else {
			 w.WriteHeader(http.StatusMethodNotAllowed)
			 w.Write([]byte("Method not allowed"))
		 }
	})

	// catalog list
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

	// prduct list
	mux.HandleFunc("GET /catalog/{id}/produto", func(w http.ResponseWriter, r *http.Request){
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