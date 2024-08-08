// db/repositoryCatalog.go
package db

import (
	"fmt"
	//"golang.org/x/text/message/catalog"
)

// "database/sql"
// "log"

type Catalog struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Imglink     string `json:"imgLink"`
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

func AddCatalog(catalog Catalog) {
	//fmt.Println(catalog)
	query := "INSERT INTO catalogs (title, description, imglink) VALUES ($1, $2, $3) RETURNING id"
	id := 0
	err := DB.QueryRow(query, catalog.Title, catalog.Description, catalog.Imglink).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New ID:", id)
}

func AlterCatalog(catalog Catalog){
	query := "UPDATE catalogs SET title = $2, description = $3, imglink = $4 WHERE id = $1;"
	res, err := DB.Exec(query, catalog.Id, catalog.Title, catalog.Description, catalog.Imglink)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Rows changed: ", count, " id: ", catalog.Id) 
}

func DeleteCatalog(id string){
	query := "DELETE FROM catalogs where id = $1;"
	res, err := DB.Exec(query, id)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Rows changed: ", count, " id: ", id) 
}
