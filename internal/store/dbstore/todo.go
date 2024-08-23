package dbstore

import (
	"goth/internal/store"

	"gorm.io/gorm"
)

type TodoStore struct {
	db *gorm.DB
}

type NewTodoStoreParams struct {
	DB *gorm.DB
}

func NewTodoStore(params NewTodoStoreParams) *TodoStore {
	return &TodoStore{
		db: params.DB,
	}
}

func (s *TodoStore) CreateTodo(title string, userID uint) (*store.Todo, error) {
	todo := &store.Todo{
		Title:     title,
		UserID:    userID,
		Completed: false,
	}
	result := s.db.Create(todo)

	if result.Error != nil {
		return nil, result.Error
	}
	return todo, nil
}

func (s *TodoStore) GetTodos(userID uint) ([]store.Todo, error) {
	var todos []store.Todo

	err := s.db.Preload("User").Where("user_id = ?", userID).Find(&todos).Error

	if err != nil || len(todos) == 0 {
		return nil, err
	}

	return todos, nil
}

func (s *TodoStore) DeleteTodo(todoID uint) error {
	var todo store.Todo

	err := s.db.Preload("User").Where("id = ?", todoID).First(&todo).Error

	if err != nil {
		return err
	}

	return s.db.Delete(&todo).Error
}
