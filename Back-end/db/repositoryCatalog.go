// db/repositoryCatalog.go
package db

import (
	"fmt"

	//"golang.org/x/text/message/catalog"
)

// "database/sql"
// "log"

type Catalog struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Imglink string `json:"imgLink"`
}

func GetCatalogs() ([]Catalog, error) {
	query := "SELECT * FROM catalogs"
    rows, err := DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var catalogs []Catalog
    for rows.Next() {
        var catalog Catalog
        if err := rows.Scan(&catalog.Id, &catalog.Title, &catalog.Description, &catalog.Imglink); err != nil {
            return nil, err
        }
        catalogs = append(catalogs, catalog)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return catalogs, nil
}

func AddCatalog(catalog Catalog){
	query := "INSERT INTO users (title, description, imglink) VALUES ($1, $2, $3) RETURNING id`"
	id := 0
	err := DB.QueryRow(query, catalog.Title, catalog.Description, catalog.Id).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New ID:", id)
}
