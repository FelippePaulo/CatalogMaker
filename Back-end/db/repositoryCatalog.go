// db/repositoryCatalog.go
package db


// "database/sql"
// "log"

type Catalog struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Imglink string `json:"imgLink"`
}

func GetCatalogs() ([]Catalog, error) {
	
    rows, err := DB.Query("SELECT * FROM catalogs")
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

//import "fmt"

//func getAllCatalogs(catalog) {
	//var catalog Catalog

	// sql := "SELECT * FROM Catalogs"
	// //fmt.Printf(sql)
	// row := db.QueryRow(sql)
	// if err := row.Scan(&catalog.Id, catalog.Title, catalog.Description, catalog.Imglink); err != nil{
	// 	if err == sql.ErrNoRows{
	// 		return catalog, fmt.Errorf("fail to load catalogs")
	// 	}
	// }

	// sqlStatement := "SELECT * FROM Catalogs"
	// id := 0
	// err = db.QueryRow(sqlStatement).Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("New record ID is:", id)



	// rows, err := db.DB.Query("SELECT * FROM Catalogs")
	// if err != nil {
	// 	log.Fatalf("Query failed: %v", err)
	// }
	// defer rows.Close()

	// for rows.Next() {
		
	// 	var catalog Catalog
	// 	if err := rows.Scan(&catalog.Id, &catalog.Title, &catalog.Description, &catalog.Imglink ); err != nil {
	// 		log.Fatalf("Row scan failed: %v", err)
	// 	}
	// 	fmt.Printf("ID: %d, Title: %s, Description: %s, ImgLink: %s\n", catalog.Id, catalog.Title, catalog.Description, catalog.Imglink)
	// }


//}