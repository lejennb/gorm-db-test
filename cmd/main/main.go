package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "log"
)

type Author struct { // authors
    ID        uint `gorm:"primaryKey"`
    Firstname string
    Lastname  string
}

type Book struct { // "books"
    ID     uint `gorm:"primaryKey"`
    Title  string
    Year   uint
    Authors []Author `gorm:"many2many:book_authors;"` // N:M books authors
}


func connectToMariaDB() (*gorm.DB, error) {
    dsn := "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}

func main() {
    /* Connect */
    db, err := connectToMariaDB()
    if err != nil {
        log.Fatal(err)
    }
    conn, err := db.DB() 
    defer conn.Close()

    /* Migrate */
    err = db.AutoMigrate(&Author{}, &Book{})
    if err != nil {
        log.Fatal(err)
    }

    /* Insert */
    authors := []Author{
        Author{ID: 1, Firstname: "Brian", Lastname: "Kernighan"},
    }
    db.Create(&authors)

    books := []Book{
        Book{ID: 1, Title: "The C Programming Language", Year: 1988, Authors: []Author{{ID: 1}}},
        Book{ID: 2, Title: "The Go Programming Language", Year: 2015, Authors: []Author{{ID: 1}}},
    }
    db.Create(&books)
}

