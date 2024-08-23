package handlers

import (
	"goth/internal/middleware"
	"goth/internal/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type DeleteTodosHandler struct {
	todoStore store.TodoStore
}

func NewDeleteTodosHandler(params TodosHandlerParams) *DeleteTodosHandler {
	return &DeleteTodosHandler{
		todoStore: params.TodoStore,
	}
}

func (h *DeleteTodosHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.UserKey).(*store.User)

	if !ok {
		http.Error(w, "Error rendering template", http.StatusForbidden)
		return
	}

	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)

	err := h.todoStore.DeleteTodo(uint(id))

	if err != nil {
		http.Error(w, "Error deleting todo", http.StatusInternalServerError)
		return
	}
}
