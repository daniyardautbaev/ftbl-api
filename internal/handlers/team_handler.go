package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/daniyardautbaev/ftbl-api/internal/db"
    "github.com/daniyardautbaev/ftbl-api/internal/models"
)

func CreateTeam(w http.ResponseWriter, r *http.Request) {
    var t models.Team
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := db.DB.Create(&t).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(t)
}

func ListTeams(w http.ResponseWriter, r *http.Request) {
    var teams []models.Team
    db.DB.Preload("Players").Find(&teams)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(teams)
}
