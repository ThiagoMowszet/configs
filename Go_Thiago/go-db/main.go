package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

type Product struct {
    Name    string
    Price   float64
    Avaible bool
}


func main() {
    // Create the connection
    connStr := "postgres://postgres:secret@localhost:5432/gopgtest?sslmode=disable"
    // Open the DB
    db, err := sql.Open("postgres", connStr)

    // defer function to close the DB
    defer db.Close()

    // Validations
    if err != nil {
        log.Fatal(err) 
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

    createProductTable(db)

    /*     product := Product{"book",15.55, true}

    pk := insertProduct(db, product)
    // fmt.Printf("ID = %d", pk)

    var name string
    var price float64
    var avaible bool

    query := `SELECT name, price, avaible FROM product WHERE id = $1`
    err2 := db.QueryRow(query, pk).Scan(&name, &price, &avaible)

    if err2 != nil {
        if err2 != sql.ErrNoRows {
            log.Fatalf("no rows found with ID %d", pk)
        }
        log.Fatal(err) 

        fmt.Printf("Name = %s\n", name)
        fmt.Printf("Price = %f\n", price)
        fmt.Printf("Avaible = %t\n", avaible)
    } */

    data := []Product{}
    rows, err := db.Query("SELECT name, price, avaible FROM product")

    defer rows.Close()

    var name string
    var price float64
    var avaible bool

    for rows.Next() {
        err := rows.Scan(&name, &price, &avaible)
        if err != nil {
            log.Fatal(err)
        }
        data = append(data, Product{name, price, avaible})
    }
    fmt.Println(data)
}

func createProductTable(db *sql.DB){
    query := `CREATE TABLE IF NOT EXISTS product(
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        price NUMERIC(6,2) NOT NULL,
        avaible BOOLEAN,
        created timestamp DEFAULT NOW()
    )`

    _, err := db.Exec(query)

    if err != nil {
        log.Fatal(err) 
    }
}

func insertProduct(db *sql.DB, product Product) int {
    query := `INSERT INTO product (name, price, avaible) 
    VALUES ($1, $2, $3) RETURNING id`

    var pk int
    err := db.QueryRow(query, product.Name, product.Price, product.Avaible).Scan(&pk)

    if err != nil {
        log.Fatal(err) 
    }

    return pk
}
