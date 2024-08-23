package handlers

import (
	"goth/internal/middleware"
	"goth/internal/store"
	"goth/internal/templates"
	"net/http"
)

type TodosHandler struct {
	todoStore store.TodoStore
}

type TodosHandlerParams struct {
	TodoStore store.TodoStore
}

func NewTodosHandler(params TodosHandlerParams) *TodosHandler {
	return &TodosHandler{
		todoStore: params.TodoStore,
	}
}

func (h *TodosHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value(middleware.UserKey).(*store.User)

	if !ok {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	todos, err := h.todoStore.GetTodos(user.ID)

	// if err != nil {
	// 	http.Error(w, "Error getting todos", http.StatusInternalServerError)
	// 	return
	// }

	c := templates.Todos(todos)
	err = templates.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
