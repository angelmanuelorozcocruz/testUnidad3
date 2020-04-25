package main

import (
	/*Librerias Importantes*/
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*Variables globales*/
var db *gorm.DB

/*Modelado estructurado del libro */
type Book struct {
	Id      	int
	Nombre  	string `gorm:"type:varchar(100);"`
	Desc 		string `gorm:"type:varchar(450);"`
	Autor  		string `gorm:"type:varchar(200);"`
	Edit   		string `gorm:"type:varchar(255);"`
	Fecha   	string `gorm:"type:varchar(255);"`
}

/*Funcion principal donde ocurre la magia*/
func main() {
	r := gin.Default()
	r.GET("/", homePage)
	r.GET("/books", getbooks)
	r.GET("/books/:id", getonebook)
	r.POST("/books", newbook)
	r.DELETE("/books/:id", deleteonebook)
	r.PUT("/books/:id", updateonebook)
	r.Run(":8080")
	fmt.Println("Open localhost:8080 in your browser")

	db, err = gorm.Open("mysql", "root:/*MyUser*/.@(127.0.0.1)/books?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

/*Funciones*/
func homePage(c *gin.Context) {
	var books []Book
	books, _ = getbooks(books)
}

func getbooks(c *gin.Context) {
	var books []Book
	var err error
	books, err = getbooks(books)	
}

func getonebook(c *gin.Context) {
	var book Book
	var autor Autor
	var err error
	id := c.Param("id")
	book, err = getonebooks(idInt, book)
}

func deleteonebook(c *gin.Context) {
	var book Book
	var err error
	errMessage := ""
	id := c.Param("id")

	if err == nil {
		if book.ID != 0 {
			err = deleteBook(book)
			if err == nil {
				c.JSON(200, gin.H{
					"message": "Eliminado correctamente.",
				})
			} else {
				errMessage = err.Error()
			}
		errMessage = err.Error()
	}
	if len(errMessage) != 0 {
		c.JSON(400, gin.H{
			"message": "Ah ocurrido un error: " + err.Error() + ".",
		})
	}
}

func updateonebook(c *gin.Context) {
	var book Book
	var err error
	id := c.Param("id")
	nombre := c.PostForm("nombre")
	des := c.PostForm("des")
	autor := c.PostForm("autor")
	edit:= c.PostForm("edit")
	fechap:= c.PostForm("fechap")
	book, err = getBookByID(idInt, book)
	if err == nil {
		if book.ID != 0 {
			if len(nombre) != 0 {
				book.Nombre = nombre
			}
			if len(des) != 0 {
				book.Des = des
			}
			if len(autor) != 0 {
				book.Autor = autor
			}
			if len(edit) != 0 {
				book.Edit = edit
			}
			if len(fechap) != 0 {
				book.Fechap = fechap
			}
			fmt.Println(book)
			err = updateBook(book)
			if err == nil {
				c.JSON(200, gin.H{
					"message": "Actualizado correctamente.",
				})
			} else {
				errMessage = err.Error()
			}
		}
	}
}

func newbook(c *gin.Context) {
	var err error
	nombre := c.PostForm("nombre")
	des := c.PostForm("des")
	autor := c.PostForm("autor")
	edit := c.PostForm("edit")
	fechap := c.PostForm("fechap")
	if len(nombre) != 0 || len(des) != 0 || len(autor) != 0 || len(edit) != 0 || len(fechap) != 0 {
		if !match {
			c.JSON(400, gin.H{
				"message": "Formato de fecha inválido. Use (DD/MM/YYYY o DD-MM-YYYY)",
			})
		} else {
			book := Book{Nombre: nombre, Des: des, Autor: autor, Edit: editorial, Fehcap: fechap}
			err = createBook(book)
			if err == nil {
				c.JSON(200, gin.H{
					"message": "Registrado correctamente",
				})
			} else {
				c.JSON(400, gin.H{
					"message": "Ocurrió un error: " + err.Error(),
				})
			}
		}
	}
}
