package main

import (
	"github.com/SlenderLizard/go-todo/handlers"
	"github.com/SlenderLizard/go-todo/initializers"
	"github.com/SlenderLizard/go-todo/middleware"
	"github.com/SlenderLizard/go-todo/repository"
	"github.com/gin-gonic/gin"
)

func init() {
	//Runs before main()
	initializers.LoadEnvVariables()

}

func main() {
	authStore := repository.NewUserStore() // In-memory user store
	todoStore := repository.NewTodoStore()

	//Handlers
	todoHandler := &handlers.TodoHandler{Repo: todoStore}
	authHandler := &handlers.AuthHandler{Repo: authStore}

	//Router
	router := gin.Default()
	router.POST("/register", authHandler.Register) //User registration,Username must be unique
	router.POST("/login", authHandler.Login)       //User login

	//Protected routes
	todoRoutes := router.Group("/todos")
	todoRoutes.Use(middleware.AuthMiddleware)
	{
		todoRoutes.GET("", todoHandler.GetTodos)          //Get all todos for authenticated user
		todoRoutes.GET("/:id", todoHandler.GetTodo)       //Get a single todo by ID for authenticated user
		todoRoutes.POST("", todoHandler.CreateTodo)       //Create a new todo for authenticated user
		todoRoutes.PUT("/:id", todoHandler.UpdateTodo)    //Update a todo by ID for authenticated user
		todoRoutes.DELETE("/:id", todoHandler.DeleteTodo) //Delete a todo by ID for authenticated user
	}

	router.Run()
}
