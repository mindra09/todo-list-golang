package repository

import (
	"errors"
	"todo-app/internal/domain"
)

type TodoRepository interface {
	FindAll() []domain.Todo
	Save(todo domain.Todo) domain.Todo
	DeleteByID(id int) error
}

type InMemoryTodoRepo struct {
	todos  []domain.Todo
	nextID int
}

func NewInMemoryTodoRepo() *InMemoryTodoRepo {
	return &InMemoryTodoRepo{
		todos:  []domain.Todo{},
		nextID: 1,
	}
}

func (r *InMemoryTodoRepo) FindAll() []domain.Todo {
	return r.todos
}

func (r *InMemoryTodoRepo) Save(todo domain.Todo) domain.Todo {
	todo.ID = r.nextID
	r.nextID++
	r.todos = append(r.todos, todo)
	return todo
}

func (r *InMemoryTodoRepo) DeleteByID(id int) error {
	for i, t := range r.todos {
		if t.ID == id {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
