package main
import (
    "log"
    "net/http"
    "os"
    "time"

    "github.com/joho/godotenv"
    "github.com/daniyardautbaev/ftbl-api/internal/db"
    "github.com/daniyardautbaev/ftbl-api/internal/models"
    "github.com/daniyardautbaev/ftbl-api/internal/router"
)

func main() {
    // load .env
    _ = godotenv.Load()

    if err := db.Connect(); err != nil {
        log.Fatal("db.Connect:", err)
    }

    // auto-migrate tables
    if err := db.DB.AutoMigrate(&models.Team{}, &models.Player{}); err != nil {
        log.Fatal("AutoMigrate:", err)
    }

    srv := &http.Server{
        Addr:         ":" + getEnv("PORT", "8080"),
        Handler:      router.NewRouter(),
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  60 * time.Second,
    }

    log.Println("Server started on", srv.Addr)
    if err := srv.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}