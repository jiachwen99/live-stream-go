package api

import "net/http"

type APIController struct {
	storage *apiStorage
}

func (controller *APIController) Create(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for creating a user
}

func (controller *APIController) GetAll(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for getting all users
}

func (controller *APIController) GetSingle(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for getting a single user
}

func (controller *APIController) Update(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for updating a user
}

func (controller *APIController) Delete(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for deleting a user
}
