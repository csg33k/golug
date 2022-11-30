package www

import (
	"github.com/safaci2000/golug/services"
	"net/http"
)

// ListDistros
// @Description List all users
// @Produce json
// @Success 200 {object} []string
// @Router /api/v1/distros/list [get]
func ListDistros(w http.ResponseWriter, r *http.Request) {
	srv := services.GetServices()
	distros, err := srv.ListDistributions()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondwithJSON(w, http.StatusOK, distros)
}
