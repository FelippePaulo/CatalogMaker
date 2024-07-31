package db

import (
   "database/sql"
   "fmt"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "felippe"
  password = "batata12"
  dbname   = "Catalogs-db"
)

// Declare a global variable to hold the database connection
var DB *sql.DB

func Initialize() error {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    
    var err error
    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        return err
    }
    
    if err = DB.Ping(); err != nil {
        return err
    }
    
    return nil
}

// func dbConecction()  {
//   psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//     "password=%s dbname=%s sslmode=disable",
//     host, port, user, password, dbname)
//   db, err := sql.Open("postgres", psqlInfo)
//   if err != nil {
//     panic(err)
//   }
//   defer db.Close()
  

//   sqlStatement := `
// INSERT INTO users (age, email, first_name, last_name)
// VALUES ($1, $2, $3, $4)
// RETURNING id`
//   id := 0
//   err = db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
//   if err != nil {
//     panic(err)
//   }
//   fmt.Println("New record ID is:", id)
//}