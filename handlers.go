package main

import (
	"dashboard-api/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome To VOE Dashboard!")
}
func (a *App) UserLogin(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {

		respondWithError(w, http.StatusBadRequest, "Invalid payload; Required [username]")
		return
	}
	defer r.Body.Close()

	user, err := models.Login(a.DB, user.Username)
	if err != nil {
		sendErrorResponse(w, err, "User not found")
		return
	}
	res := map[string]models.User{}
	res["data"] = user
	respondWithJSON(w, http.StatusOK, res)

}
func (a *App) UsersIndex(w http.ResponseWriter, r *http.Request) {
	users, err := models.UserList(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	res := map[string]models.Users{}
	res["data"] = users
	respondWithJSON(w, http.StatusOK, res)
}

func (a *App) UserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request Data: User ID")
		return
	}
	user, err := models.UserByID(a.DB, id)
	if err != nil {
		sendErrorResponse(w, err, "User not found")
		return
	}
	res := map[string]models.User{}
	res["data"] = user
	respondWithJSON(w, http.StatusOK, res)

}

func (a *App) TeamIndex(w http.ResponseWriter, r *http.Request) {
	teams, err := models.TeamList(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	res := map[string]models.Teams{}
	res["data"] = teams
	respondWithJSON(w, http.StatusOK, res)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)

	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(code)
	w.Write(response)
}

func sendErrorResponse(w http.ResponseWriter, err error, msg string) {
	switch err {
	case sql.ErrNoRows:
		respondWithError(w, http.StatusNotFound, msg)
	default:
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	return
}
