package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-app/internal/usecase"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	uc *usecase.TodoUsecase
}

func NewTodoHandler(uc *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{uc: uc}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title string `json:"title"`
		File  string `json:"file"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	todo := h.uc.CreateTodo(req.Title, req.File)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.uc.GetTodos())
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "param must int", http.StatusBadRequest)
		return
	}

	if err := h.uc.DeleteTodo(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
