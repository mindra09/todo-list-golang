package usecase

import (
	"todo-app/internal/domain"
	"todo-app/internal/repository"
	"todo-app/internal/service"
)

type TodoUsecase struct {
	repo repository.TodoRepository
	ai   service.AIService
}

func NewTodoUsecase(repo repository.TodoRepository, ai service.AIService) *TodoUsecase {
	return &TodoUsecase{
		repo: repo,
		ai:   ai,
	}
}

func (u *TodoUsecase) GetTodos() []domain.Todo {
	return u.repo.FindAll()
}

func (u *TodoUsecase) CreateTodo(title, file string) domain.Todo {
	category := "General"

	if file != "" {
		if aiCat, err := u.ai.PredictCategory(file); err == nil {
			category = aiCat
		}
	}

	todo := domain.Todo{
		Title:      title,
		Completed:  false,
		AICategory: category,
		Category:   category,
		Refined:    false,
	}

	return u.repo.Save(todo)

}

func (u *TodoUsecase) DeleteTodo(id int) error {
	return u.repo.DeleteByID(id)
}
