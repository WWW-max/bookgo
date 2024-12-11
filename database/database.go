package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// init DB connector
func InitDatabase() {
	var err error

	dbUser := "root"
	dbPass := "123456"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "library"
    
	// mysql DSN (Data Source Name)
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	    dbUser, dbPass, dbHost, dbPort, dbName,
    )
    
	// connect to DB
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to MySQL database: ", err)
	}

	// test connection
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping MySQL database: ", err)
	}

	log.Println("Connected to MySQL database!")
    
	// init tables
	initTables();
}

// init tables
func initTables() {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS books (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			author VARCHAR(255) NOT NULL,
			isbn VARCHAR(20) NOT NULL
		);
	`
	_, err := DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create books table: ", err)
	}
}