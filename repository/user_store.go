package repository

import (
	"errors"
	"sync"

	"github.com/SlenderLizard/go-todo/models"
)

var ErrUserExists = errors.New("username already exists")

type UserStore struct {
	mu     sync.RWMutex
	users  map[string]models.User
	nextID int
}

// NewUserStore initializes and returns a new UserStore
func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[string]models.User),
		nextID: 1,
	}
}

// Create adds a new user to the store
func (store *UserStore) Create(user models.User) error {
	store.mu.Lock()
	defer store.mu.Unlock()
	if _, exists := store.users[user.Username]; exists {
		return ErrUserExists
	}

	user.ID = store.nextID
	store.nextID++

	store.users[user.Username] = user
	return nil
}

// GetByUsername retrieves a user by their username
func (store *UserStore) GetByUsername(username string) (models.User, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	user, exists := store.users[username]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}
