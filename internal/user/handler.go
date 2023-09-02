package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Handler is a collection of all the service handlers.
type Handler struct {
	UserService Service
}

// NewHandler creates a new Handler.
func NewHandler(us Service) *Handler {
	return &Handler{
		UserService: us,
	}
}

// Create creates a new user.
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.UserService.Create(&u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%+v", u)
}

// Get gets a user by ID.
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	u, err := h.UserService.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%+v", u)
}

// GetAll gets all users.
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	us, err := h.UserService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%+v", us)
}

// Update updates a user.
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.UserService.Update(&u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%+v", u)
}

// Delete deletes a user.
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if err := h.UserService.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Deleted user with ID: %d", id)
}
