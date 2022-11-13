package www

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/safaci2000/golug/models"
	"github.com/safaci2000/golug/services"
	"net/http"
	"strconv"
)

// ListUsers
// @Description List all users
// @Produce json
// @Success 200 {object} models.LinuxUser
// @Router /api/v1/users/list [get]
func ListUsers(w http.ResponseWriter, r *http.Request) {
	srv := services.GetServices()
	users := srv.ListUsers()
	respondwithJSON(w, http.StatusOK, users)
}

// CreateUser
// @Description Create a New User
// @Produce json
// @Param request body models.LinuxUser true "query params"
// @Success 200 {object} models.LinuxUser
// @Router /api/v1/users/ [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var newUser models.LinuxUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid Request received")
	}
	srv := services.GetServices()
	createdUser, err := srv.CreateUser(newUser)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not save user to DB")
	}
	respondwithJSON(w, http.StatusOK, createdUser)
}

// Update
// @Description Create a New User
// @Produce json
// @Param id path int true "user id"
// @Param request body models.LinuxUser true "query params"
// @Success 200 {object} models.LinuxUser
// @Router /api/v1/users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var newUser models.LinuxUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	userId := extractPathUserId(w, r)
	if err != nil || newUser.Id != userId {
		respondWithError(w, http.StatusInternalServerError, "Invalid Request received")
	}
	srv := services.GetServices()
	updatedUser, err := srv.UpdateUser(newUser)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not save user to DB")
	}
	respondwithJSON(w, http.StatusOK, updatedUser)
}

// DeleteUser
// @Description Retrieves a User by ID
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} models.LinuxUser
// @Router /api/v1/users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	rawUserId := chi.URLParam(r, "id") // 👈 getting path param
	svc := services.GetServices()
	userId, err := strconv.ParseInt(rawUserId, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid userID")
	}
	err = svc.DeleteUser(userId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not delete user")
	}

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "User Successfully deleted", "error": ""})
}

func extractPathUserId(w http.ResponseWriter, r *http.Request) int64 {
	rawUserId := chi.URLParam(r, "id") // 👈 getting path param
	userId, err := strconv.ParseInt(rawUserId, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid userID")
	}
	return userId

}

// GetUser
// @Description Retrieves a User by ID
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} models.LinuxUser
// @Router /api/v1/users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := extractPathUserId(w, r)
	svc := services.GetServices()
	user, err := svc.GetUser(userId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid userID")
	}
	respondwithJSON(w, http.StatusOK, user)
}
