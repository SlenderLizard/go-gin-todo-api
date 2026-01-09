package handlers

import (
	"net/http"
	"strconv"

	"github.com/SlenderLizard/go-todo/models"
	"github.com/SlenderLizard/go-todo/repository"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Repo *repository.TodoStore
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	user := c.GetInt("userID")
	list := h.Repo.GetAll(user)
	c.JSON(http.StatusAccepted, list)
}

func (h *TodoHandler) GetTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format in URL"})
		return
	}

	todo, exists := h.Repo.GetByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	//Check ownership of the todo
	userID := c.GetInt("userID")
	if todo.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to view this todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var userInput struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
	}

	err := c.BindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	todo := models.Todo{
		Title:       userInput.Title,
		Description: userInput.Description,
		Completed:   userInput.Completed,
		UserID:      c.GetInt("userID"),
	}

	//Create todo in the store
	createdTodo := h.Repo.Create(todo)
	c.JSON(http.StatusCreated, createdTodo)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format in URL"})
		return
	}

	var userInput struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Completed   *bool   `json:"completed"`
	}

	err = c.BindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	todoRepo, exists := h.Repo.GetByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	//Check ownership of the todo
	userID := c.GetInt("userID")
	if todoRepo.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this todo"})
		return
	}

	//Only update fields that are not nil in userInput

	if userInput.Title != nil {
		todoRepo.Title = *userInput.Title
	}
	if userInput.Description != nil {
		todoRepo.Description = *userInput.Description
	}
	if userInput.Completed != nil {
		todoRepo.Completed = *userInput.Completed
	}

	//Update todo in the store
	updatedTodo, ok := h.Repo.Update(id, todoRepo)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format in URL"})
		return
	}

	todo, exists := h.Repo.GetByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	//Check ownership of the todo
	userID := c.GetInt("userID")
	if todo.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this todo"})
		return
	}

	//Delete todo from the store
	ok := h.Repo.Delete(id)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
