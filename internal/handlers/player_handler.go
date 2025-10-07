package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "github.com/daniyardautbaev/ftbl-api/internal/db"
    "github.com/daniyardautbaev/ftbl-api/internal/models"
)

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
    var p models.Player
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := db.DB.Create(&p).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(p)
}

func ListPlayers(w http.ResponseWriter, r *http.Request) {
    var players []models.Player
    db.DB.Find(&players)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(players)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, _ := strconv.Atoi(idStr)
    var p models.Player
    if err := db.DB.First(&p, id).Error; err != nil {
        http.Error(w, "Player not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(p)
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, _ := strconv.Atoi(idStr)
    var p models.Player
    if err := db.DB.First(&p, id).Error; err != nil {
        http.Error(w, "Player not found", http.StatusNotFound)
        return
    }
    var input models.Player
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    p.Name = input.Name
    p.Position = input.Position
    p.Goals = input.Goals
    p.Assists = input.Assists
    p.Matches = input.Matches
    p.TeamID = input.TeamID

    if err := db.DB.Save(&p).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(p)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, _ := strconv.Atoi(idStr)
    if err := db.DB.Delete(&models.Player{}, id).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
