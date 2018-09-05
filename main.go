package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type PESSOA struct {
	IDPessoa  uint   `gorm:"primary_key" json:"idpessoa,omitempty" `
	Nome      string `gorm:"size:255" json:"nome,omitempty" `
	Sobrenome string `gorm:"size:255" json:"sobrenome,omitempty" `
	Apelido   string `gorm:"size:255" json:"apelido,omitempty" `
	Likes     int    `gorm:"size:11" json:"likes,omitempty" `
	Deslikes  int    `gorm:"size:11" json:"deslikes,omitempty" `
}

var (
	server   = "localhost"
	port     = 1480
	user     = "sa"
	password = "@Password123"
	database = "LIKEPEOPLE"
)

func main() {

	router := gin.Default()

	// Adiciona o FrontEnd Views na raiz da rota
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		// Exibit todas as pessoas cadastradas
		api.GET("/pessoas", func(c *gin.Context) {
			c.JSON(http.StatusOK, getPessoas())
		})

		// Exibir uma pessoa utilizando o ID como identificador
		api.GET("/pessoas/:id", func(c *gin.Context) {
			id := c.Params.ByName("id")
			user_id, _ := strconv.Atoi(id)
			c.JSON(http.StatusOK, getPessoasId(user_id))
		})

		// Criar Pessoa
		api.POST("/pessoas", func(c *gin.Context) {
			var novoUsuario PESSOA
			c.ShouldBindJSON(&novoUsuario)
			criarPessoa(novoUsuario)
			c.JSON(http.StatusOK, "OK")
		})

		// Exibir uma pessoa utilizando o ID como identificador
		api.DELETE("/pessoas/:id", func(c *gin.Context) {
			id := c.Params.ByName("id")
			user_id, _ := strconv.Atoi(id)
			c.JSON(http.StatusOK, deletePessoaId(user_id))
		})

		api.PUT("/pessoas", func(c *gin.Context) {
			var attUsuario PESSOA
			c.ShouldBindJSON(&attUsuario)
			criarPessoa(attUsuario)
			c.JSON(http.StatusOK, "OK")
		})
	}

	router.Run(":3005")
}

func getConnection() (db *gorm.DB) {

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)

	db, err := gorm.Open("mssql", connectionString)

	if err != nil {
		panic("failed to connect database")
		fmt.Printf(err.Error())
		return
	}

	db.AutoMigrate(&PESSOA{})

	return db
}

func criarPessoa(pessoa PESSOA) {
	// Cria conexão com banco de dados
	db := getConnection()
	// Cria nova pessoa no DB
	db.Create(&pessoa)
	// Fecha conexão com o banco de dados
	db.Close()
}

func getPessoas() (pessoas []PESSOA) {

	// Array de pessoas
	var tdPessoas []PESSOA

	// Cria conexão com o banco de dados
	db := getConnection()

	// Busca todos usuários e os adiciona ao array tdPessoas
	db.Find(&tdPessoas)

	// Encerra conexão com o banco de dados
	db.Close()

	// Retorna array de pessoas
	return tdPessoas
}

func getPessoasId(id int) (pessoa PESSOA) {
	var returnPessoa PESSOA

	db := getConnection()
	db.Find(&returnPessoa, id)
	db.Close()
	return returnPessoa
}

func deletePessoaId(id int) PESSOA {
	var returnPessoa PESSOA

	db := getConnection()
	db.Delete(&returnPessoa, id)
	db.Close()
	return returnPessoa
}

func putPessoa(pessoa PESSOA) PESSOA {
	var attPessoa PESSOA
	// Cria conexão com banco de dados
	db := getConnection()
	// Cria nova pessoa no DB
	db.Find(&attPessoa, 15)
	db.Save(&attPessoa)
	// Fecha conexão com o banco de dados
	db.Close()
	fmt.Print(pessoa)
	return attPessoa
}
