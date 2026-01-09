package handlers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/SlenderLizard/go-todo/models"
	"github.com/SlenderLizard/go-todo/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Repo *repository.UserStore
}

// Register handles user registration
func (ah *AuthHandler) Register(c *gin.Context) {
	var userInput struct {
		Username string `json:"username" binding:"required,min=3,max=25"`
		Password string `json:"password" binding:"required,min=8"`
	}

	err := c.BindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	//Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Username: userInput.Username,
		Password: string(hash),
	}

	err = ah.Repo.Create(user)
	if err != nil {
		if errors.Is(err, repository.ErrUserExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login handles user login
func (ah *AuthHandler) Login(c *gin.Context) {
	var userInput struct {
		Username string `json:"username" binding:"required,min=3,max=25"`
		Password string `json:"password" binding:"required,min=8"`
	}

	err := c.BindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	//Retrieve user from the store
	user, err := ah.Repo.GetByUsername(userInput.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if bcryptErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}


	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
