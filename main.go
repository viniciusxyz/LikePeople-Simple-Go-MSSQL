package main

import (
	"fmt"
	"log"
	"context"
	"net/http"
	"database/sql"
	"github.com/gin-gonic/contrib/static"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

var server = "127.0.0.1"
var port = 1480
var user = "sa"
var password = "@Password123"
var database = "LIKEPEOPLE"

func main() {

	/*
		Conexão com banco de dados
	*/

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
	server, user, password, port, database)

	var err error

    // Create connection pool
    db, err = sql.Open("sqlserver", connString)
    if err != nil {
        log.Fatal("Error creating connection pool:", err.Error())
    }
	fmt.Printf("Connected!\n")
	
	count, err := ReadEmployees()
    fmt.Printf("Read %d rows successfully.\n", count)

	router := gin.Default()

	// Adiciona o FrontEnd Views na raiz da rota
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/pessoas", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/teste", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Teste Group API GO",
			})
		})
	}

	router.Run(":3002")
}


func ReadEmployees() (int, error) {

    ctx := context.Background()

    // Check if database is alive.
    err := db.PingContext(ctx)
    if err != nil {
        log.Fatal("Error pinging database: " + err.Error())
    }

    tsql := fmt.Sprintf("SELECT ID_PESSOA,NOME,APELIDO from PESSOAS")

    // Execute query
    rows, err := db.QueryContext(ctx, tsql)
    if err != nil {
        log.Fatal("Error reading rows: " + err.Error())
        return -1, err
    }

    defer rows.Close()

    var count int = 0

    // Iterate through the result set.
    for rows.Next() {
        var NOME, APELIDO string
        var ID_PESSOA int

        // Get values from row.
        err := rows.Scan(&ID_PESSOA, &NOME, &APELIDO)
        if err != nil {
            log.Fatal("Error reading rows: " + err.Error())
            return -1, err
        }

        fmt.Printf("ID Pessoa: %d, Nome: %s, Apelido: %s\n", ID_PESSOA, NOME, APELIDO)
        count++
    }

    return count, nil
}