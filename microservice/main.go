package main

import (
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()
	tt := router.Group("/api/v1/todos")
	{
		tt.POST("/", createTodo)
		tt.GET("/", fetchAllTodo)
		tt.GET("/:id", fetchSingleTodo)
		tt.PUT("/:id", updateTodo)
		tt.DELETE("/:id", deleteTodo)

	}

	router.Run()
}

func createTodo(c *gin.Context) {
	completed, _ := strconv.ParseBool(c.PostForm("completed"))
	todo := todoModel{
		Title:     c.PostForm("title"),
		Completed: completed,
	}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceID": todo.ID})
}
func fetchAllTodo(c *gin.Context) {
	var todos []todoModel

	db.Find(&todos)
	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todos})
}
func fetchSingleTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found"})
		return
	}
	t := todoModel{
		Title:     todo.Title,
		Completed: todo.Completed,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": t})
}
func updateTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found"})
		return
	}
	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}
func deleteTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found"})
		return
	}
	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully"})
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&todoModel{}); err != nil {
		log.Fatal(err)
	}
}

type todoModel struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
