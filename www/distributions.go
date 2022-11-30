package www

import (
	"github.com/safaci2000/golug/services"
	"net/http"
)

// ListDistros
// @Description List all users
// @Produce json
// @Success 200 {object} []string
// @Router /api/v1/distro/list [get]
func ListDistros(w http.ResponseWriter, r *http.Request) {
	srv := services.GetServices()
	distros, err := srv.ListDistributions()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondwithJSON(w, http.StatusOK, distros)
}

// ListDistros
// @Description List all users
// @Produce json
// @Success 200 {object} map[string]int64
// @Router /api/v1/distro/count [get]
func ListDistroCount(w http.ResponseWriter, r *http.Request) {
	srv := services.GetServices()
	distros, err := srv.LinuxDistroCount()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	result := make(map[string]int64)
	for _, val := range distros {
		result[val.LinuxDistro] = val.Count
	}
	respondwithJSON(w, http.StatusOK, result)
}
