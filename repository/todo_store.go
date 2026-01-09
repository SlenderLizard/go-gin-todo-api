package repository

import (
	"sync"

	"github.com/SlenderLizard/go-todo/models"
)

type TodoStore struct {
	mu     sync.RWMutex        //mutex for concurrent access
	todos  map[int]models.Todo //map of ID to Todo, would create problems if becomes too large
	nextID int                 //ID increases only in create, and do not decrease
}

// NewTodoStore initializes and returns a new TodoStore
func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos:  make(map[int]models.Todo),
		nextID: 1,
	}
}

// GetAll returns all todos in the store, sonra bunu tekli yaparız
func (store *TodoStore) GetAll(id int) []models.Todo {
	store.mu.RLock()
	defer store.mu.RUnlock()

	todos := make([]models.Todo, 0, len(store.todos))
	for _, todo := range store.todos {
		//Check ownership, then add
		if todo.UserID != id {
			continue
		}
		todos = append(todos, todo)
	}
	return todos
}

// GetByID retrieves a todo by its ID
func (store *TodoStore) GetByID(id int) (models.Todo, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	todo, exists := store.todos[id]
	return todo, exists
}

// Create adds a new todo to the store
func (store *TodoStore) Create(todo models.Todo) models.Todo {
	store.mu.Lock()
	defer store.mu.Unlock()

	todo.ID = store.nextID
	store.todos[todo.ID] = todo
	store.nextID++
	return todo
}

// Update modifies an existing todo in the store
func (store *TodoStore) Update(id int, updatedTodo models.Todo) (models.Todo, bool) {
	store.mu.Lock()
	defer store.mu.Unlock()

	_, exists := store.todos[id]
	if !exists {
		return models.Todo{}, false
	}

	store.todos[id] = updatedTodo

	return updatedTodo, true
}

// Delete removes a todo from the store by its ID
func (store *TodoStore) Delete(id int) bool {
	store.mu.Lock()
	defer store.mu.Unlock()

	_, exists := store.todos[id]
	if exists {
		delete(store.todos, id)
	}
	return exists
}
