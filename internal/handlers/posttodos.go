package handlers

import (
	"goth/internal/middleware"
	"goth/internal/store"
	"goth/internal/templates"
	"net/http"
)

type PostTodosHandler struct {
	todoStore store.TodoStore
}

func NewPostTodosHandler(params TodosHandlerParams) *PostTodosHandler {
	return &PostTodosHandler{
		todoStore: params.TodoStore,
	}
}

func (h *PostTodosHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middleware.UserKey).(*store.User)

	if !ok {
		http.Error(w, "Error rendering template", http.StatusForbidden)
		return
	}

	title := r.FormValue("title")
	todo, _ := h.todoStore.CreateTodo(title, user.ID)

	c := templates.PostTodoSuccess(*todo)
	err := c.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}

}
